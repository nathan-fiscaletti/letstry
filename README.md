```markdown
# letstry

letstry is a powerful tool designed to streamline project creation and management within VSCode, built in Golang. It provides users with a suite of features to create, manage, and export project templates, ensuring a smooth workflow for developers.

## Features

- **Project Creation**: Create new projects with a simple command. The tool initializes a new directory and opens VSCode within that directory.
- **Template Management**: Save, export, import, list, and remove project templates with ease. Templates can be tagged with languages for better organization.
- **Session Tracking**: Track active VSCode sessions by PID, and prompt users to save their projects upon closing the VSCode window.
- **Export and Save**: Export projects to a specified path or save them as templates for future use.
- **Custom Templates**: Set a default template, update it, or create new ones based on your needs.

## Installation

To install letstry, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/letstry.git
    ```
2. Navigate to the project directory:
    ```sh
    cd letstry
    ```
3. Build the project:
    ```sh
    go build
    ```
4. Move the binary to your desired location and ensure it's in your PATH.

## Usage

### Creating a New Project

To create a new project, use the following command:
```sh
lt new
```
You can specify a template and a name for the session:
```sh
lt new -template <template_name> -name <session_name>
```
If no name is specified, the PID will be used as the session name. If no template is specified, the default template will be used.

### Managing Templates

Set the default template:
```sh
lt template default set <template_name>
```
Check the current default template:
```sh
lt template default
```
Export a template:
```sh
lt template export <template_name> <path>
```
Import a template:
```sh
lt template import <path>
```
List all templates:
```sh
lt template list
```
Remove a template:
```sh
lt template remove <template_name>
```
Tag a template with a language:
```sh
lt template tag <template_name> <language>
```
Remove a tag from a template:
```sh
lt template untag <template_name> <language>
```
List all tags for a template:
```sh
lt template tags <template_name>
```
List all templates for a specific language:
```sh
lt template list <language>
```

### Managing Sessions

List all active sessions:
```sh
lt session list
```
Kill a session:
```sh
lt session kill <name|pid>
```
Export a session:
```sh
lt session export <name|pid> <path>
```
Save a session as a template:
```sh
lt session save <name|pid> <template_name>
```

## Workflow

1. **Creating a New Project**: When you create a new project with `lt new`, a directory is created, and VSCode opens in that directory. The project can be based on a specified template.
2. **Tracking Sessions**: Each VSCode process is tracked by PID in a `~/.letstry` file. This allows the tool to monitor processes and prompt users to save their work before closing.
3. **Managing Templates**: Templates are an essential part of letstry, enabling you to quickly set up projects with predefined structures. You can manage these templates using the provided commands.

## Contributing

We welcome contributions to improve letstry. If you have suggestions or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
```

Feel free to customize this README further to suit your specific project requirements!