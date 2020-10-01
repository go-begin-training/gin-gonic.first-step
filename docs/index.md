## Welcome to GitHub Pages

You can use the [editor on GitHub](https://github.com/go-begin-training/gin-gonic.first-step/edit/main/docs/index.md) to maintain and preview the content for your website in Markdown files.

Whenever you commit to this repository, GitHub Pages will run [Jekyll](https://jekyllrb.com/) to rebuild the pages in your site, from the content in your Markdown files.

### Markdown

Markdown is a lightweight and easy-to-use syntax for styling your writing. It includes conventions for

# server gin-gonic (first step - get acquainted with REST API")
server with gin-gonic framework (level 1)

### How can this project run? ###

We use the files "go.mod" and "go.sum" to contain the necessary configuration packages.
If the project you pull requires importing packages, open a terminal at the project root directory and enter the following command:
```
  $ go mod download
```

Next, run directly with the command:
```
  $ go run * .go
```
Or build into a program with the command:
```
  $ CGO_ENABLE=0 go build --ldflags "-extldflags \"-static\"-s -w" -o bin/application -trimpath ./*.go
```

The program after being built will be saved in `-o bin/application`. Request the system to execute with the command:
```
  $ /bin/bash -c bin/application
```

Documents:
- https://gin-gonic.com/
- https://gin-gonic.com/docs/
- https://godoc.org/github.com/gin-gonic/gin
- https://github.com/gin-gonic/gin
- https://github.com/gin-gonic/gin#api-examples

---

For more details see [GitHub Flavored Markdown](https://guides.github.com/features/mastering-markdown/).

### Jekyll Themes

Your Pages site will use the layout and styles from the Jekyll theme you have selected in your [repository settings](https://github.com/go-begin-training/gin-gonic.first-step/settings). The name of this theme is saved in the Jekyll `_config.yml` configuration file.

### Support or Contact

Having trouble with Pages? Check out our [documentation](https://docs.github.com/categories/github-pages-basics/) or [contact support](https://github.com/contact) and we’ll help you sort it out.
