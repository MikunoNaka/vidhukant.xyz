#!/bin/sh

cmd="sass $@"

$cmd web/styles/styles.scss:../blog/css/styles.css &
$cmd web/styles/home.scss:../blog/css/home.css &
$cmd web/styles/post/posts.scss:../blog/css/posts.css &
$cmd web/styles/post/post.scss:../blog/css/post.css &

