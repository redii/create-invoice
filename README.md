# create-invoice
[![Build Status](https://travis-ci.org/redii/create-invoice.svg?branch=master)](https://travis-ci.org/redii/create-invoice)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## 🧾 About this app..
A simple tool which generates basic pdf invoice documents. The application has a predefined layout in which you can fill your own data by customizing the `template.txt` and `data.txt` files. When executing the tool, it will read these input files and build an invoice pdf document.

## ⚙ How to use
I used the [gofpdf library](https://github.com/jung-kurt/gofpdf) to build this tool. The Layout is defined by multiple different cell objects within the pdf (see sample below), which are filled with the data provided by the input files. 
The input elements in these files are divided from one another by `\n;\n`, and so can be easily customized with any basic text editor.

The `template.txt` file contains the inputs, which should be very consistant and mostly never change for a single user while  the `data.txt` file contains all inputs for each individual invoice like reference number and date. When you decide to change and run the code by yourself you can toggle the `devMode` const, which enables the borders around each cell in the output file.

## 📌 Sample
![pdf screen](https://i.imgur.com/k7xpamx.png) ![pdf screen dev](https://i.imgur.com/4RciqHN.png)

## ✔ Todo
- [ ] multiline table cells
