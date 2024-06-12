# letstry: Use Case — Temporarily Working on a Project

When you need to work on a project temporarily and don't want to clutter your workspace, letstry offers a solution to create a temporary project space. This use case demonstrates how to create a temporary project, manage your work, and save your progress for future reference. Additionally, it covers creating a new project based on a template or a Git repository.

## Creating a Temporary Project

To create a temporary project, use the following command:
```sh
lt new -template <template_name> -name <session_name>
```
- **template_name**: The name of the template to use for the project. If not specified, the default template is used.
- **session_name**: The name of the session. If not specified, the PID of the VSCode process is used.

Example:
```sh
lt new -template web-app -name my-temp-project
```

This command initializes a temporary project directory based on the specified template and opens it in VSCode. You can start working on your project immediately.

## Managing Your Work

While you are working on your temporary project, letstry tracks the VSCode session using the PID. If you close the VSCode window and you haven't configured letstry to display the "Do you want to save your project" dialog, the temporary directory will be deleted. Therefore, it's crucial to save your work before closing VSCode.

## Saving Your Progress

If you decide that you want to keep the work you've done in your temporary project, you have several options to save your progress:

### Exporting the Project

To export the project to a more permanent location on your file system, use:
```sh
lt session export <name|pid> <path>
```
Example:
```sh
lt session export my-temp-project /path/to/save
```

### Saving as a Template

To save the project as a template for future use, use:
```sh
lt session save <name|pid> <template_name>
```
Example:
```sh
lt session save my-temp-project my-new-template
```

### Exporting to a Git Repository

To export the project to a Git repository, use:
```sh
lt session export <name|pid> <repository_url>
```
Example:
```sh
lt session export my-temp-project https://github.com/yourusername/my-temp-project.git
```

When exporting to a Git repository, you also have the option to automatically delete the files after the repository has been created. This helps keep your workspace clean and uncluttered.

## Summary

Using letstry to create and manage temporary projects allows you to quickly prototype ideas, test frameworks, or perform tasks without cluttering your workspace. By exporting or saving your progress, you ensure that your work is preserved and can be continued or reused in the future.

For more information on other use cases, refer to the [letstry Use Cases](../use-cases.md) document.