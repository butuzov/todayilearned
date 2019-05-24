# Visual Studio Code & Extensions Tips

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

* [Ui & Editing](#ui-editing)
	* [Keyboard Shortcuts: Editing (Custom)](#keyboard-shortcuts-editing-custom)
	* [Keyboard Shortcuts: Multiline Cursor (MLC) Editor](#keyboard-shortcuts-multiline-cursor-mlc-editor)
	* [Keyboard Shortcuts: Panels and Sidebars (Custom)](#keyboard-shortcuts-panels-and-sidebars-custom)
* [Extension: Markdown Oreview Enhanced](#extension-markdown-oreview-enhanced)
		* [Tips & Tricks](#tips-tricks)
		* [Keyboard Shortcuts (Custom)](#keyboard-shortcuts-custom)
		* [Markdown Styling](#markdown-styling)

<!-- /code_chunk_output -->

## Ui & Editing

### Keyboard Shortcuts: Editing (Custom)
|        Shortcut                                       | Description
|-------------------------------------------------------|-------------------------
| <kbd>⇧</kbd> <kbd>^</kbd> <kbd>D</kbd>                | Copy Line Up
| <kbd>⌘</kbd> <kbd>/</kbd>                             | Toggle comments

### Keyboard Shortcuts: Multiline Cursor (MLC) Editor 
|        Shortcut                                          | Description
|----------------------------------------------------------|-------------------
| <kbd>⇧</kbd> <kbd>⌥</kbd> <kbd>⌘</kbd> <kbd>↓</kbd>      | Column Selection
| <kbd>⇧</kbd> <kbd>⌘</kbd> <kbd>L</kbd>                   | Search Results to MLC

### Keyboard Shortcuts: Panels and Sidebars (Custom)
|        Shortcut                                            | Description
|------------------------------------------------------------|-------------------
| <kbd>⌘</kbd> <kbd>B</kbd>                                  | Open in split window
| <kbd>⌘</kbd> <kbd>'</kbd>                                  | Toggle Console Position
| <kbd>⌘</kbd> <kbd>J</kbd>                                  | Toggle Console Visibility
| <kbd>⌘</kbd> <kbd>\\</kbd>                                 | Toggle Sidebar Visibility
| <kbd>⌘</kbd> <kbd>0-9</kbd>                                | Switch to edit window


## Extension: Markdown Oreview Enhanced

#### Tips & Tricks

 * Insert TOC for Quick TOC Generation

#### Keyboard Shortcuts (Custom)

|  Keys                                | Commands                                | Description
|--------------------------------------|-----------------------------------------|-------------------
| <kbd>^</kbd><kbd>M</kbd>             | `markdown-preview-enhanced.openPreview` | Open Preview


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
