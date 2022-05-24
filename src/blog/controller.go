/*
 * vidhukant.xyz: Vidhu Kant's Personal Website
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package blog

import (
  "strconv"
  "fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
  ID        int
  CreatedAt string
  UpdatedAt *string
  Title     string
  Content   string
  Tags      []string
}

func (db *dbhandler) getPostCount(tag *int) int {
  var query string
  if tag != nil {
    query = "SELECT COUNT(DISTINCT(PostID)) FROM Post_Tags WHERE TagID = " + strconv.Itoa(*tag)
  } else {
	// because some posts might not even have tags
    query = "SELECT COUNT(*) FROM Posts"
  }
  fmt.Println("qry: ", query)

  rows, err := db.connection.Prepare(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var count int
  if err := rows.QueryRow().Scan(&count); err != nil {
    panic(err)
  }

  return count
}

// start = read from nth row, limit = read n rows
func (db *dbhandler) getPosts(start, limit int) []Post {
  rows, err := db.connection.Query(
    fmt.Sprintf("SELECT ID, DATE_FORMAT(CreatedAt, \"%%D %%M %%Y\"), Title FROM Posts LIMIT %d,%d", start, limit),
  ) 
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var posts []Post
  for rows.Next() {
    var p Post
    err := rows.Scan(&p.ID, &p.CreatedAt, &p.Title)
    if err != nil {
      panic(err)
    }
    // load post's tags
    p.Tags = db.getPostTags(p.ID)
    posts = append(posts, p)
  }

  return posts
}

// same as getPosts but in descending order
func (db *dbhandler) getPostsReverse(start, limit int) []Post {
  rows, err := db.connection.Query(
    fmt.Sprintf("SELECT ID, DATE_FORMAT(CreatedAt, \"%%D %%M %%Y\"), Title FROM Posts ORDER BY ID DESC LIMIT %d,%d", start, limit),
  ) 
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var posts []Post
  for rows.Next() {
    var p Post
    err := rows.Scan(&p.ID, &p.CreatedAt, &p.Title)
    if err != nil {
      panic(err)
    }
    // load post's tags
    p.Tags = db.getPostTags(p.ID)
    posts = append(posts, p)
  }

  return posts
}

func (db *dbhandler) getPost(id int) Post {
  rows, err := db.connection.Prepare(
    `SELECT ID, DATE_FORMAT(CreatedAt, "%D %M %Y"), DATE_FORMAT(UpdatedAt, "%D %M %Y"), Title, Content FROM Posts WHERE ID = ?`,
  )
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  var post Post
  if err := rows.QueryRow(id).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title, &post.Content); err != nil {
    // TODO: handle error when rows are empty
    post.Content = "404"
    // panic(err)
  }

  return post
}

func (db *dbhandler) getPostTags(id int) []string {

  rows, err := db.connection.Query(
    `SELECT Tags.Name FROM Post_Tags 
    INNER JOIN Tags ON Post_Tags.TagID = Tags.ID 
    WHERE Post_Tags.PostID = ` + strconv.Itoa(id),
  ) 
  if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

  var tags []string
  for rows.Next() {
    var tag string
    err := rows.Scan(&tag)
    if err != nil {
      panic(err)
    }
    tags = append(tags, tag)
  }

  return tags
}
