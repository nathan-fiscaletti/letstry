# letstry: Use Case — Testing a Framework or Library

When you need to quickly test a new framework or library, letstry provides a streamlined way to set up a temporary project environment. This allows you to experiment without cluttering your file system. This use case details how to create a new project, experiment with the framework or library, and manage your work efficiently.

## Creating a Temporary Project for Testing

To create a temporary project for testing a framework or library, use the following command:
```sh
lt new -template <template_name> -name <session_name>
```
- **template_name**: The name of the template to use for the project. If not specified, the default template is used.
- **session_name**: The name of the session. If not specified, the PID of the VSCode process is used.

Example:
```sh
lt new -template test-framework -name framework-test
```

This command initializes a temporary project directory based on the specified template and opens it in VSCode. You can now start experimenting with the new framework or library.

## Experimenting with the Framework or Library

Once your temporary project is set up, you can install the framework or library you want to test. Use the terminal within VSCode to install dependencies and start coding. For example, if you are testing a new JavaScript framework, you might run:
```sh
npm install <framework>
```
You can then write sample code, run tests, and evaluate the framework's features and performance within this isolated environment.

## Managing Your Work

While working on your temporary project, letstry tracks the VSCode session using the PID. If you close the VSCode window and you haven't configured letstry to display the "Do you want to save your project" dialog, the temporary directory will be deleted. To avoid losing your work, make sure to save your progress before closing VSCode.

## Saving Your Progress

If you find the framework or library useful and want to keep your test project, you have several options to save your progress:

### Exporting the Project

To export the project to a more permanent location on your file system, use:
```sh
lt session export <name|pid> <path>
```
Example:
```sh
lt session export framework-test /path/to/save
```

### Saving as a Template

To save the project as a template for future use, use:
```sh
lt session save <name|pid> <template_name>
```
Example:
```sh
lt session save framework-test my-framework-template
```

### Exporting to a Git Repository

To export the project to a Git repository, use:
```sh
lt session export <name|pid> <repository_url>
```
Example:
```sh
lt session export framework-test https://github.com/yourusername/framework-test.git
```

When exporting to a Git repository, you also have the option to automatically delete the files after the repository has been created. This helps keep your workspace clean and uncluttered.

## Summary

Using letstry to test frameworks or libraries allows you to quickly set up a temporary project environment, experiment without cluttering your file system, and save your progress if you find the framework or library useful. By exporting or saving your project, you ensure that your work is preserved and can be continued or reused in the future.

For more information on other use cases, refer to the [letstry Use Cases](../use-cases.md) document.