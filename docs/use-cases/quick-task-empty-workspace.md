# letstry: Use Case — Doing a Quick Task that Requires an Empty Workspace

Sometimes, you need a clean slate to quickly accomplish a task. letstry's default "empty" template is perfect for such scenarios. This use case covers how to create an empty project space, perform your task, and manage the temporary directory.

## Creating an Empty Project Workspace

To create an empty project workspace, use the following command:
```sh
lt new -template empty -name <session_name>
```
- **session_name**: The name of the session. If not specified, the PID of the VSCode process is used.

Example:
```sh
lt new -template empty -name quick-task
```

This command initializes a temporary project directory with an empty template and opens it in VSCode. You now have a clean workspace to perform your task.

## Performing Your Task

With the empty project workspace set up, you can focus on your quick task. Whether it's writing a script, testing a snippet of code, or any other short-term task, the temporary project ensures your main workspace remains uncluttered.

## Managing Your Work

While you are working on your quick task, letstry tracks the VSCode session using the PID. If you close the VSCode window and you haven't configured letstry to display the "Do you want to save your project" dialog, the temporary directory will be deleted. To avoid losing your work, make sure to save your progress before closing VSCode.

## Saving Your Progress

If you decide that you want to keep the work you've done in your empty project workspace, you have several options to save your progress:

### Exporting the Project

To export the project to a more permanent location on your file system, use:
```sh
lt session export <name|pid> <path>
```
Example:
```sh
lt session export quick-task /path/to/save
```

### Saving as a Template

To save the project as a template for future use, use:
```sh
lt session save <name|pid> <template_name>
```
Example:
```sh
lt session save quick-task my-empty-template
```

### Exporting to a Git Repository

To export the project to a Git repository, use:
```sh
lt session export <name|pid> <repository_url>
```
Example:
```sh
lt session export quick-task https://github.com/yourusername/quick-task.git
```

When exporting to a Git repository, you also have the option to automatically delete the files after the repository has been created. This helps keep your workspace clean and uncluttered.

## Summary

Using letstry for quick tasks that require an empty workspace allows you to set up a clean project environment, perform your task efficiently, and save your progress if necessary. By exporting or saving your project, you ensure that your work is preserved and can be continued or reused in the future.

For more information on other use cases, refer to the [letstry Use Cases](../use-cases.md) document.