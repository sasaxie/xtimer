# xtimer [![Build Status](https://travis-ci.org/sasaxie/xtimer.svg?branch=master)](https://travis-ci.org/sasaxie/xtimer) [![codecov](https://codecov.io/gh/sasaxie/xtimer/branch/master/graph/badge.svg)](https://codecov.io/gh/sasaxie/xtimer)
Golang timer tools

## Installation
> go get github.com/sasaxie/xtimer

## 功能
- [x] StartTimer(interval, func(time))：定时执行一次逻辑
- [x] StartTicker(interval, func(time))：定时执行多次逻辑，直到主动关闭或者程序结束。
