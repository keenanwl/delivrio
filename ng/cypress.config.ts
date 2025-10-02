import {defineConfig} from "cypress";
import {spawn} from "child_process";

// Spawn instead of exec since the stdout buffer needs to be big
// when logging to the console. May not be an issue running as --detach though
const spawnPromise = (command: string, args: string[] = []) =>
	new Promise<void>((resolve, reject) => {
		const child = spawn(command, args, {stdio: 'inherit'});

		child.on('close', (code) => {
			if (code !== 0) {
				return reject(new Error(`Command failed with exit code ${code}`));
			}
			resolve();
		});

		child.on('error', (error) => {
			console.error(`Error executing command: ${error.message}`);
			reject(error);
		});
	});

export default defineConfig({
	e2e: {
		setupNodeEvents(on, config) {
			on('task', {
				'db:teardown': async () => {
					await spawnPromise('npx e2e:stop-server');
					return true;
				},
				'db:seed': async () => {
					await spawnPromise('npm', ['run', 'e2e:start-server']);
					console.log("Started server...")
					await spawnPromise('docker', ['ps', '-a']);
					console.log("Running containers...")
					await spawnPromise('docker', ['compose', 'logs']);
					console.log("Docker logs...")
					await spawnPromise('npx', ['wait-on', '-v', 'http-get://0.0.0.0:8080/api/health-check']);
					return true;
				},
			});
		},
		baseUrl: 'http://localhost:8080',
	},
});
