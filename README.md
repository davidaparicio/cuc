# CUC - Command-line URL Checker (and notifier)

<p align="center">
<img width="285" height="285" src="assets/img/cuc.jpg" alt="CUC logo / A gopher playing with streams of data, logo, digital art, drawing" title="CUC / Generated by DALL·E" />
</p>

| Citation |
| --------:|
| In Go we trust, its power we wield, |
| A language simple, yet revealed, |
| With concurrency at its core, |
| It's efficiency we adore. |
| |
| "Clear is better than clever," |
| Golang's mantra we deliver, |
| For code that's easy to read, |
| Is the key to success, indeed. |
| Like [Go Proverbs - Simple, Poetic, Pithy](https://go-proverbs.github.io/) |

[![Go Report Card](https://goreportcard.com/badge/davidaparicio/cuc)](https://goreportcard.com/report/davidaparicio/cuc)
[![GoDoc](https://pkg.go.dev/badge/github.com/davidaparicio/cuc?status.svg)](https://pkg.go.dev/github.com/davidaparicio/cuc)
[![Github](https://img.shields.io/static/v1?label=github&logo=github&color=E24329&message=main&style=flat-square)](https://github.com/davidaparicio/cuc)
[![Docker Pulls](https://img.shields.io/docker/pulls/davidaparicio/cuc.svg)](https://hub.docker.com/r/davidaparicio/cuc)
[![Maintenance](https://img.shields.io/maintenance/yes/2024.svg)]()

<!-- [![GitLab](https://img.shields.io/static/v1?label=gitlab&logo=gitlab&color=E24329&message=main&style=flat-square)](https://gitlab.com/davidaparicio/cuc) -->
<!-- [![Froggit](https://img.shields.io/static/v1?label=froggit&logo=froggit&color=red&message=no&style=flat-square)](https://lab.frogg.it/davidaparicio/cuc) -->

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/davidaparicio/cuc/blob/master/LICENSE.md)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fdavidaparicio%2Fcuc.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fdavidaparicio%2Fcuc?ref=badge_shield)
[![Known Vulnerabilities](https://snyk.io/test/github/davidaparicio/cuc/badge.svg)](https://snyk.io/test/github/davidaparicio/cuc)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=davidaparicio_cuc&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=davidaparicio_cuc)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=davidaparicio_cuc&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=davidaparicio_cuc)
[![Twitter](https://img.shields.io/twitter/follow/dadideo.svg?style=social)](https://twitter.com/intent/follow?screen_name=dadideo)

CUC (English pronunciation: [_cuc_] / λευκός) is a very simple CLI tool to check various HTTP status for example if a webpage is available (200) or not found (404).

> It's delicious like a [TUC](https://en.wikipedia.org/wiki/TUC_(cracker)) (cracker), Biscuit of the French company LU, but with a C, for the [Console](https://en.wikipedia.org/wiki/Command-line_interface).

---

<!--
. **[Overview](#overview)** .
**[Features](#features)** .
**[Supported backends](#supported-backends)** .
**[Quickstart](#quickstart)** .
**[Web UI](#web-ui)** .
**[Documentation](#documentation)** .

. **[Support](#support)** .
**[Release cycle](#release-cycle)** .
**[Contributing](#contributing)** .
**[Maintainers](#maintainers)** .
**[Credits](#credits)** .
-->

. **[Overview](#overview)** .
**[Quickstart](#quickstart)** .
**[Credits](#credits)** .

---

## Overview

## Quickstart

<!-- 
If you have already ```Docker``` installed on your laptop

```docker run davidaparicio/cuc:<TAG/VERSION_LIKE_v0.0.5> -u <WEBSITE_TO_CHECK> -c 200 -f <PATH_TO_AUDIO_FILE>```

If not,
-->

You need ```Go``` and all [dependencies](https://deps.dev/go/github.com%2Fdavidaparicio%2Fcuc/v0.0.0-20230313221521-d867c87d3847/dependencies)

```CGO_ENABLED=1 go run ./main.go https://www.example.com/ loop --URL -s 10 -c 200 -f assets/mp3/ubuntu_dialog_info.mp3```

For more information, you can see [examples here](EXAMPLES.md)

## Credits

"A gopher playing with streams of data, logo, digital art, drawing" generated by <a href="https://labs.openai.com/" target="_blank">DALL·E (OpenAI)</a>