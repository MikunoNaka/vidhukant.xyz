#!/bin/bash

mkdir -p css/projects

cmd="sass $@"

$cmd scss/styles.scss:css/styles.css
$cmd scss/font-faces.scss:css/font-faces.css
$cmd scss/index.scss:css/index.css

$cmd scss/projects/styles.scss:css/projects/styles.css
