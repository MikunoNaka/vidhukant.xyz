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
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
  ID        int
  CreatedAt string
  UpdatedAt string
  Title     string
  Content   string
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
  }

  return post
}
