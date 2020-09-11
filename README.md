# jenkobs
Open Build Service to Jenkins action daemon via AMQP

This is a little daemon that is watching status of whatever project on OBS
and then reacts according to the configured actions. Currently supported:

- HTTP action.
- Shell action

HTTP action will call specific defined URL with passed parameters to trigger something.
This is useful for CI testing to watch when a specific package has been built

Shell action will call a command. This is useful to call e.g. a Configuration Management
system with a specific state/playbook/etc that will do the rest on an expected package or project in OBS.
