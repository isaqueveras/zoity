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
  run         run service
  services    get services
  version     

Flags:
  -h, --help   help for zoity

Use &quot;zoity [command] --help&quot; for more information about a command.
</pre>

## Add service

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity add --name powersso --command &quot;go run main.go&quot; --path ~/path/powersso
zoity:<font color="#26A269"><b> service configured successfully</b></font>
</pre>

## Run services

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity run powersso-ui powersso powersso-test
zoity:<font color="#26A269"><b> pid=140108: the powersso-ui service has been initialized</b></font>
zoity:<font color="#26A269"><b> pid=140109: the powersso service has been initialized</b></font>
zoity:<font color="#C01C28"><b> service powersso-test not found</b></font>
</pre>

## List services

<pre><font color="#26A269"><b>isaque@veras</b></font>:<font color="#12488B"><b>~</b></font>$ zoity services

| ID         | NAME                      | CREATED         | COMMAND                       |
|------------|---------------------------|-----------------|-------------------------------|
| qvate8n8   | powersso                  | 2023-11-15      | go run *.go                   |
| v0ko2o6v   | powersso-ui               | 2023-11-15      | npm run start                 |
</pre>
