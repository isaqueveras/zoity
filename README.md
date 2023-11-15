# zoity
Zoity is an orchestrator for configuring and running services locally.

## Commands

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity
Zoity is an orchestrator for configuring and running services locally.

Usage:
  zoity [command]

Available Commands:
  add         Add a service
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Use to initialize Zoity configuration
  services    get services
  version     

Flags:
  -h, --help   help for zoity

Use &quot;zoity [command] --help&quot; for more information about a command.
</pre>

## Add service

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity add --name powersso --port 4747 --command &quot;go run main.go&quot; --path ~/path/powersso
zoity: service configured successfully
</pre>

## List services

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity services

| ID         | NAME                      | PORT       | CREATED         | COMMAND                       |
|------------|---------------------------|------------|-----------------|-------------------------------|
| qvate8n8   | powersso                  | :4747      | 2023-11-15      | go run *.go                   |
| v0ko2o6v   | powersso-ui               | :3030      | 2023-11-15      | npm run start                 |
</pre>
