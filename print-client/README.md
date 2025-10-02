# DELIVRIO Print Client
When DELIVRIO core is deployed in server mode, this app serves as the workstation client for picking up print jobs.

## Running the Application in Developer Mode
The easiest way is to use the Wails CLI: `wails dev`

This should hot refresh when making changes the Frontend and rebuild when making changes in the Go.

## Building the Application for Production
When building for Production, there are some parts to notice:
 - Angular is going to use the `frontend/src/environments/environment.prod.ts` when building the frontend.
 - You should properly change the Icons for your Application in `build/` and `build/windows` (when building for Windows)
 - The Developer Console will not work.

Finally, you can build you Application with: `wails build`
