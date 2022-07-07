#!/bin/bash

cmd="sass $@"

$cmd scss/font-faces.scss:css/font-faces.css
$cmd scss/styles.scss:css/styles.css
