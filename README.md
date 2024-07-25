# letstry

letstry is a powerful tool designed to streamline project creation and management within VSCode, built in Golang. It provides users with a suite of features to create, manage, and export project templates, ensuring a smooth workflow for developers.

> If you want to understand in what ways letstry can help you, you are highly encouraged to read the [Use Cases](./docs/use-cases.md) section.

## Features

- **Project Creation**: Create new projects with a simple command. The tool initializes a temporary directory and opens it in VSCode for quick prototyping.
- **Template Management**: Save, export, import, list, and remove project templates with ease. Templates can be tagged with languages for better organization.
- **Session Tracking**: Track active VSCode sessions by PID, and prompt users to save their projects upon closing the VSCode window.
- **Export and Save**: Export projects to a specified path, save them as templates for future use, or export them to a git repository.
- **Custom Templates**: Set a default template, update it, or create new ones based on your needs.

## Installation

To install letstry, run the following command:

```sh
$ go install github.com/nathan-fiscaletti/letstry/cmd/letstry@latest
```

## Usage

### Project Creation

Creating a new project with letstry is simple and efficient. Use the `lt new` command to initialize a temporary project directory and open it in VSCode. This allows for quick prototyping. If you like the results, you can export the project to a more permanent location or save it as a template. 

```sh
lt new -template <template_name> -name <session_name>
```

If the VSCode window is closed and you haven't configured letstry to display the "Do you want to save your project" dialog, the temporary directory will be deleted. Therefore, you should either export your project using `lt session export` or save it as a template using `lt session save`.

### Managing Templates

Templates in letstry allow you to save and reuse project structures. You can manage templates using various commands:

- **Set Default Template**: Update the default template used for new projects.
- **Export and Import Templates**: Share templates with others or use them on different machines.
- **List and Remove Templates**: Keep your template library organized.
- **Tagging Templates**: Tag templates with languages to quickly find the right template for your project.

These capabilities help maintain consistency and efficiency, especially when working on multiple projects or collaborating with others.

### Session Tracking

letstry tracks active VSCode sessions by their PID. This tracking allows the tool to prompt you to save your work when you close VSCode, preventing data loss. Commands related to sessions include listing active sessions, killing a session, exporting a session, and saving a session as a template.

```sh
lt session list
lt session kill <name|pid>
lt session export <name|pid> <path>
lt session save <name|pid> <template_name>
```

Session tracking ensures that you always have control over your active projects and can easily manage them.

### Export and Save

The export and save features allow you to export projects to a specified path, save them as templates for future use, or export them to a git repository. This functionality is particularly useful for creating backups or sharing your project setup with others.

```sh
lt template export <template_name> <path>
lt template import <path>
lt session export <name|pid> <path>
lt session save <name|pid> <template_name>
lt session export <name|pid> <repository_url>
```

By using these commands, you can ensure that your work is always preserved and easily transferable.

## Workflow

1. **Creating a New Project**: When you create a new project with `lt new`, a temporary directory is created, and VSCode opens in that directory. The project can be based on a specified template.
2. **Tracking Sessions**: Each VSCode process is tracked by PID in a `~/.letstry` file. This allows the tool to monitor processes and prompt users to save their work before closing. If not saved, the temporary directory is deleted upon closing VSCode.
3. **Managing Templates**: Templates are an essential part of letstry, enabling you to quickly set up projects with predefined structures. You can manage these templates using the provided commands.

## Contributing

We welcome contributions to improve letstry. If you have suggestions or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.