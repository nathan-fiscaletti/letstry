# letstry: Use Case — Prototyping a New Idea

For developers who often have new ideas to prototype, letstry provides a streamlined way to start working on these ideas immediately. This use case explains how to create a new project for prototyping, save your progress, and export your prototype if you decide to continue working on it.

## Creating a New Project for Prototyping

To create a new project for prototyping a new idea, use the following command:
```sh
lt new -template <template_name> -name <session_name>
```
- **template_name**: The name of the template to use for the project. If not specified, the default template is used.
- **session_name**: The name of the session. If not specified, the PID of the VSCode process is used.

Example:
```sh
lt new -template prototype -name new-idea
```

This command initializes a temporary project directory based on the specified template and opens it in VSCode. You can start working on your new idea immediately.

## Developing Your Prototype

With your temporary project set up, you can focus on developing your prototype. Write code, test features, and iterate on your idea within this isolated environment. The temporary project setup ensures that your main workspace remains uncluttered.

## Managing Your Work

While you are working on your prototype, letstry tracks the VSCode session using the PID. If you close the VSCode window and you haven't configured letstry to display the "Do you want to save your project" dialog, the temporary directory will be deleted. To avoid losing your work, make sure to save your progress before closing VSCode.

## Saving Your Progress

If you decide that your prototype is worth keeping, you have several options to save your progress:

### Exporting the Project

To export the project to a more permanent location on your file system, use:
```sh
lt session export <name|pid> <path>
```
Example:
```sh
lt session export new-idea /path/to/save
```

### Saving as a Template

To save the project as a template for future use, use:
```sh
lt session save <name|pid> <template_name>
```
Example:
```sh
lt session save new-idea my-prototype-template
```

### Exporting to a Git Repository

To export the project to a Git repository, use:
```sh
lt session export <name|pid> <repository_url>
```
Example:
```sh
lt session export new-idea https://github.com/yourusername/new-idea.git
```

When exporting to a Git repository, you also have the option to automatically delete the files after the repository has been created. This helps keep your workspace clean and uncluttered.

## Summary

Using letstry to prototype new ideas allows you to quickly set up a project environment, focus on development without cluttering your main workspace, and save your progress if the prototype proves valuable. By exporting or saving your project, you ensure that your work is preserved and can be continued or reused in the future.

For more information on other use cases, refer to the [letstry Use Cases](../use-cases.md) document.