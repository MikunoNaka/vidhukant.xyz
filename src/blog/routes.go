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
	"html/template"
	"net/http"
	"strconv"

	"github.com/MikunoNaka/vidhukant.xyz/db"
	"github.com/gin-gonic/gin"
)

// database connection
var base *dbhandler
func init() {
  connection := db.ConnectDB()
  base = newHandler(connection)
}

func getPosts(ctx *gin.Context) {
  limitOptions := []int{10, 20, 30}
  limit := 10
  // if limit is in url query use that
  if l := ctx.Query("limit"); l != "" {
    limit, _ = strconv.Atoi(l)
  }

  pageNum := 1
  // if pageNum is in url query use that
  if p := ctx.Query("page"); p != "" {
    pageNum, _ = strconv.Atoi(p)
    // pageNum can't be less than 1
    if pageNum < 1 { pageNum = 1 }
  }

  // if firstPost is in url query use that
  firstPost := limit * (pageNum - 1)
  if f := ctx.Query("first"); f != "" {
    firstPost, _ = strconv.Atoi(f)
    // firstPost can't be less than 0
    if firstPost < 0 {firstPost = 0}
  }

  // get posts from database
  posts := base.getPosts(firstPost, limit)

  showNext := true
  // TODO: if tags are specified, replace with nil
  // check if difference between all post count and posts shown is same
  if base.getPostCount(nil) - (firstPost + len(posts)) < 1 {
    showNext = false
  }

  // TODO: calculate if nextpage/prevpage button should be hidden
  ctx.HTML(http.StatusOK, "views/posts.html", gin.H {
    "LimitOptions": limitOptions,
    "Limit": limit,
    "FirstPost": firstPost,
    "PageNumber": pageNum,
    "PrevPage": pageNum - 1,
	  "ShowPrev": !(firstPost == 0),
	  "ShowNext": showNext,
    "NextPage": pageNum + 1,
    "PrevFirst": firstPost - limit,
    "NextFirst": firstPost + limit,
    "Posts": posts,
  })
}

func getPost(ctx *gin.Context) {
  id, _ := strconv.Atoi(ctx.Param("id"))
  post := base.getPost(id)

  ctx.HTML(http.StatusOK, "views/post.html", gin.H {
    "Title": post.Title,
    "CreatedAt": post.CreatedAt,
    "UpdatedAt": post.UpdatedAt,
    "Content": template.HTML(post.Content),
  })
}

func HomePage(ctx *gin.Context) {
  recentPosts := base.getPostsReverse(0, 10)
  // TODO: also render some tags

  ctx.HTML(http.StatusOK, "views/home.html", gin.H {
    "RecentPosts": recentPosts,
  })
}
