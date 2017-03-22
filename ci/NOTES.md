# Concourse Notes

## Plugin: Git Metadata

Store information about the current Git commit. If we can, maybe also store things like a changelog.
This may be possible by being able to see the previous version.

## Plugin: Job Status

Use job level `on_success` and `on_failure` to `put` their respective status, with the default 
status set to `error` or something using a `get` in the `plan`. Then in `ensure` we can use the 
written status in other things, for example, Slack notifications (perhaps via a custom handler).
