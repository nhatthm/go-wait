# Wait

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/go-wait)](https://github.com/nhatthm/go-wait/releases/latest)
[![Build Status](https://github.com/nhatthm/go-wait/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/go-wait/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/go-wait/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/go-wait)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/wait)
[![Donate](https://img.shields.io/badge/%20-Donate-%20?style=flat&logo=githubsponsors&color=E5E4E2)](http://donate.nhat.me)

A simple library to wait for something.

## Prerequisites

- `Go >= 1.23`

## Install

```bash
go get go.nhat.io/wait
```

## Usage

```go
package main

import (
	"context"
	"time"

	"go.nhat.io/wait"
)

func execute(ctx context.Context) error {
	if err := wait.ForDuration(time.Minute).Wait(ctx); err != nil {
        	return err
	}

	// Do something here.

	return nil
}
```

## Donation

If this project saved you some development time, buy me a cup of coffee :)

[![donate](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](http://donate.nhat.me)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://github.com/nhatthm/donate.nhat.me/blob/master/images/qr_sponsor.png" width="147px" />

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)
