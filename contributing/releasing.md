# Releasing the SDK

Releases to [`grafana-foundation-sdk`](https://github.com/grafana/grafana-foundation-sdk/)
are automated using a [few bash scripts](../scripts/) and [GitHub Actions](../.github/workflows/):

* `.github/workflows/release.yaml` is responsible for preparing and performing releases
* `.github/workflows/docs.yaml` builds and publishes the documentation
* `.github/workflows/[language]-release.yaml` perform language-specific releases

## The release process

Releasing a new version of the SDK is an automated process and follows these steps:

1. push commits to `main` (ie: merge PRs to main)
    * `.github/workflows/release.yaml` starts, and calls the [`scripts/prepare-release.sh`](../scripts/prepare-release.sh) script
    * this script runs the codegen pipeline, and publishes the results in a `release-preview` branch for which a PR is opened with the title "Next release"
1. repeat step 1 until a new SDK release is required
1. merge the "Next release" PR
    * `.github/workflows/release.yaml` starts, and calls the [`scripts/release.sh`](../scripts/release.sh) script
    * a tag is published
1. documentation and language-specific workflows start automatically whenever a tag is pushed

> [!NOTE]
> The `./scripts/prepare-release.sh` and `./scripts/release.sh` scripts are
> also usable locally.

## Dashboard Converter

Releases are triggered by pushing a tag with the format `dashboard-converter-vX.Y.Z`. The [Dashboard Converter workflow](https://github.com/grafana/grafana-foundation-sdk/actions/workflows/dashboard-converter.yaml) runs automatically on tag push.

```console
git tag dashboard-converter-v1.0.0
git push origin dashboard-converter-v1.0.0
```
