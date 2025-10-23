# go-fiber-exercise
https://go-fiber-app-smy2gd5xla-uw.a.run.app/

A minimal web api  built with the Fiber framework and deployed automatically to Google Cloud Run through a GitHub Actions CI/CD pipeline.

Make any change in the repository and push to main. This will trigger the automated pipeline.

When deployed, the app exposes a single HTTP endpoint (`/`) that returns a compact JSON object with:
```json
{
   "deployed_at": "10-22-2025"
  "message": "My name is Nico Reiss",
  "timestamp": 1730000000000
}
