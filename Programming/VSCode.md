# Visual Studio Code & Extensions Tips

## Ui & Editing

### Keyboard Shortcuts: Editing (Custom)
|        Shortcut                                         | ID                                     | Description
|---------------------------------------------------------|----------------------------------------|--------------
| <kbd>⌘</kbd> <kbd>/</kbd>                               | `editor.action.commentLine`            | Toggle Line Comment(s)
| <kbd>⌘</kbd> <kbd>[</kbd>                               | `editor.action.outdentLines`           | Outdent Line(s) / Selection
| <kbd>⌘</kbd> <kbd>]</kbd>                               | `editor.action.indentLines`            | Indent Line(s) / Selection
| <kbd>⌘</kbd> <kbd>k</kbd> `+` <kbd>⌘</kbd> <kbd>h</kbd> | `editor.foldAll`                       | Fold All
| <kbd>⌘</kbd> <kbd>k</kbd> `+` <kbd>⌘</kbd> <kbd>j</kbd> | `editor.unfoldAll`                     | Unfold All
| <kbd>⌘</kbd> <kbd>k</kbd> `+` <kbd>⌘</kbd> <kbd>c</kbd> | `editor.foldAllBlockComments`          | Fold Block Comments
| <kbd>^</kbd> <kbd>⌥</kbd> <kbd>↑</kbd>                  | `cursorPageUp`                         | Cursor Page Up
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>⌥</kbd> <kbd>↑</kbd>     | `cursorPageUpSelect`                   | Cursor Page Up + Selection
| <kbd>^</kbd> <kbd>⌥</kbd> <kbd>↓</kbd>                  | `cursorPageDown`                       | Cursor Page Down
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>⌥</kbd> <kbd>↓</kbd>     | `cursorPageDownSelect`                 | Cursor Page Down + Selection
| <kbd>⌥</kbd> <kbd>⌘</kbd> <kbd>↑</kbd>                  | `editor.action.insertCursorAbove`      | Cursor (Insert) Above
| <kbd>⌥</kbd> <kbd>⌘</kbd> <kbd>↓</kbd>                  | `editor.action.insertCursorBelow`      | Cursor (Insert) Below
| <kbd>⌥</kbd> <kbd>⇧</kbd> <kbd>↑</kbd>                  | `editor.action.copyLinesUpAction`      | Copy Line Up
| <kbd>⌥</kbd> <kbd>⇧</kbd> <kbd>↓</kbd>                  | `editor.action.copyLinesDownAction`    | Copy Line Down
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>↑</kbd>                  | `editor.action.smartSelect.expand`     | Smart Selection (Expand)
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>↓<kbd>                   | `editor.action.smartSelect.shrink`     | Smart Selection (Shrink)
| <kbd>^</kbd> <kbd>a</kbd>                               | `cursorLineStart`                      | Cursor at Line Start
| <kbd>^</kbd> <kbd>e</kbd>                               | `cursorLineEnd`                        | Cursor at Line End
| <kbd>^</kbd> <kbd>u</kbd>                               | `deleteAllLeft`                        | Delete to Left
| <kbd>^</kbd> <kbd>k</kbd>                               | `deleteAllRight`                       | Delete to Right
| <kbd>⌘</kbd> <kbd>↩</kbd>                               | `editor.action.insertLineAfter`        | Insert Line Above
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>↩</kbd>                  | `editor.action.insertLineBefore`       | Insert Line Below
| <kbd>⌥</kbd> <kbd>↑</kbd>                               | `editor.action.moveLinesUpAction`      | Move Line(s)/Selection Up
| <kbd>⌥</kbd> <kbd>↓</kbd>                               | `editor.action.moveLinesDownAction`    | Move Line(s)/Selection Down
| <kbd>⌥</kbd> <kbd>↓</kbd>                               | `editor.action.moveLinesDownAction`    | Move Line(s)/Selection Down

### Keyboard Shortcuts: Refactoring

|        Shortcut                                           | ID                                     | Description
|-----------------------------------------------------------|----------------------------------------|--------------
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>r</kbd>                    | `editor.action.refactor`               | Refactor
| <kbd>⌘</kbd> <kbd>e</kbd>                                 | `editor.action.rename`                 | Rename Symbol

### Keyboard Shortcuts: Navigation/Search

|        Shortcut                                           | ID                                     | Description
|-----------------------------------------------------------|----------------------------------------|--------------
| <kbd>^</kbd> <kbd>g</kbd>                                 | `workbench.action.gotoLine`            | Go To Line Number
| <kbd>⌘</kbd> <kbd>g</kbd>                                 | `editor.action.nextMatchFindAction`    | Find Next
| <kbd>⌘</kbd> <kbd>g</kbd> <kbd>⇧</kbd>                    | `editor.action.previousMatchFindAction`| Find Previous

### Keyboard Shortcuts: Multiline Cursor (MLC) Editor

|        Shortcut                                          | ID                          | Description
|----------------------------------------------------------|-----------------------------|----------------------
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>⌥</kbd> <kbd>↓</kbd>      | `cursorColumnSelectDown`    | Column Selection
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>L</kbd>                   | `addCursorsAtSearchResults` | Search Results to Multiline Cursor

### Keyboard Shortcuts: Panels and Sidebars (Custom)

|        Shortcut                                          | ID                                         | Description
|----------------------------------------------------------|--------------------------------------------|--------------
| <kbd>^</kbd> <kbd>w</kbd>                                | `workbench.action.switchWindow`            | Switch to another VSCode Window
| <kbd>^</kbd> <kbd>⇧</kbd> <kbd>p</kbd>                   | `workbench.action.showAllEditors`          | (CurrWindow) Show / Switch In Editors
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>\\</kbd>                  | `workbench.action.toggleSidebarVisibility` | Toggle Sidebar
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>e</kbd>                   | `workbench.view.explorer`                  | Sidebar: Explorer
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>\\</kbd>                  | `workbench.action.togglePanel`             | Toggle Panel
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>o</kbd>                   | `workbench.action.output.toggleOutput`     | Panels: Output
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>p</kbd>                   | `workbench.actions.view.toggleProblems`    | Panels: Problems
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>s</kbd>                   | `workbench.view.search.focus`              | Panels: Search Focus
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>t</kbd>                   | `workbench.action.terminal.toggleTerminal` | Panels: Terminal
| <kbd>^</kbd> <kbd>⌘</kbd> <kbd>f</kbd>                   | `workbench.action.toggleFullScreen`        | Toggle Full Screen
| <kbd>^</kbd> <kbd>S</kbd> + <kbd>/</kbd>                 | `workbench.action.splitEditorRight`        | Split Editor
| <kbd>^</kbd> <kbd>S</kbd> + <kbd>→</kbd>                 | `workbench.action.splitEditorRight`        | Split Editor : Right
| <kbd>^</kbd> <kbd>S</kbd> + <kbd>←</kbd>                 | `workbench.action.splitEditorRight`        | Split Editor : Left
| <kbd>^</kbd> <kbd>S</kbd> + <kbd>↑</kbd>                 | `workbench.action.splitEditorDown`         | Split Editor : Down
| <kbd>^</kbd> <kbd>S</kbd> + <kbd>↓</kbd>                 | `workbench.action.splitEditorUp`           | Split Editor : Up
| <kbd>⌘</kbd> <kbd>0-9</kbd>                              | `workbench.action.focusFirstEditorGroup`   | Focus on 1st (or `n`st) Editor Group


## Keyboard Shortcuts (Custom)

|  Keys                                 | Commands                                | Description
|---------------------------------------|-----------------------------------------|-------------------
| <kbd>⌥</kbd> <kbd>M</kbd>             | `markdown.showPreviewToSide`            | Open Preview

## Snippets

 * https://snippet-generator.app/




#### Markdown Styling

CSS Configuration isn't saved using sync extension, Save it by running "Customize CSS" comand for MPE.

```less
/* Please visit the URL below for more information: */
/* https://shd101wyy.github.io/markdown-preview-enhanced/#/customize-css */
.markdown-preview.markdown-preview {
	width:100%;
	max-width:978px;

	a {color:#0366d6;}

	h1,h2 {
		border-bottom: 1px solid #eaecef;
		padding-bottom: .3em; }

	kbd {
		& + kbd {
			margin-left:1px;
		}
		background-color: #fafbfc;
		border: 1px solid #c6cbd1;
		border-bottom-color: #959da5;
		border-radius: 3px;
		box-shadow: inset 0 -1px 0 #959da5;
		color: #444d56;
		display: inline-block;
		font-size: 11px;
		line-height: 10px;
		padding: 3px 5px;
		vertical-align: middle;
	}

	pre, code {
		background-color: #f6f8fa;
	}
}
```


### Settings

* [`editor.cursorSurroundingLines`](https://twitter.com/njukidreborn/status/1160952980595605504) - many visible lines around the cursor you can see while moving cursor towards viewport (`@todo` check in next version).
