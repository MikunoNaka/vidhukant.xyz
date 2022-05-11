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
  posts := base.getPosts();

  ctx.HTML(http.StatusOK, "views/posts.html", gin.H {
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
