# Go App Boilerplate on App Engine Standard go 1.11 runtime

Get started with Go 1.11 Standard by reading [documentation here](https://cloud.google.com/appengine/docs/standard/go111/). You'll need to create an App Engine project of course.
Define your [app.yaml runtime settings](https://cloud.google.com/appengine/docs/standard/go111/config/appref).
Build your Create React App with `npm run build` to create the `build` directory.
Define your server and routing in your main.go file, and since we're serving up a [Create React App](https://github.com/facebook/create-react-app) front-end We're pointing the server to the `./build` directory for static files.
For deploying go 1.11, we're using the go.mod method instead of the former GOPATH method. See documentation on that [update here](https://cloud.google.com/appengine/docs/standard/go111/go-differences).
Use the following commands to create a go.mod before deployment:
```
$ export GO111MODULE=on  # required to use this module method
$ go mod init
$ go mod tidy
```
## Setting Up Firebase Firestore
For more info on how to set up your firebase service account firebase.json file, [go here](https://firebase.google.com/docs/admin/setup). You generate them here: `https://console.firebase.google.com/u/1/project/<your_firebase_project_id>/settings/serviceaccounts/adminsdk`

For use of other project [flags and services with gcloud deployments](https://cloud.google.com/sdk/gcloud/reference/services/), define those in your `deploy.sh` file. Then deploy your app with `npm run deploy`.