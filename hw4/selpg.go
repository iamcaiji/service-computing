package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"github.com/spf13/pflag"
)

type selgpNode struct {
	start 		int    
	end   		int    
	pagelen    	int   
	pagetype   	string 
	inputfile  	string 
	dest 		string 
	other_ele 	[]string
}

var args selgpNode

func show_detail() {
	fmt.Println("Usage: ./selpg -s[start_page_number] -e[end_page_number] [-l[page_size_number(default 72)] or -f] [input_filename]")
	fmt.Println("-e[number]	: start page.")
	fmt.Println("-s[number]	: end page.")
	fmt.Println("-l[number] : how many lines in one page.")
	fmt.Println("-l 		: 72(default) lines in one page.")
	fmt.Println("-d[des_f_n]: destination file.")
	fmt.Println("[filename] : inputfile, it is not neccssary.")
}

func main() {
	// get the args
	get(&args)
	//anylyse if the args is valid
	analyse(&args)
	//run and get what we need
	run(&args)
}

func get(args *selgpNode) {
	pflag.IntVarP(&args.start, "start", "s", 0, "start page")
	pflag.IntVarP(&args.end, "end", "e", 0, "end page")
	pflag.IntVarP(&args.pagelen, "pagelen", "l", 72, "how many lines in one page") 
	pflag.StringVarP(&args.pagetype, "pagetype", "f", "l", "...which way to split page")
	pflag.StringVarP(&args.dest, "destination", "d", "", "dest")
	pflag.Usage = show_detail
	pflag.Parse()
}

func analyse(args *selgpNode) {
	temp := pflag.Args()
	//if no input file so maybe the NO.1 arg is a number
	if len(temp) > 0 && !(temp[0][0] <= '9' && temp[0][0] >= '0'){
		args.inputfile = temp[0]
	} else {
		args.inputfile = ""
		if len(temp) > 0{
		args.other_ele = temp[:];
		}
	}
	if args.start <= 0 || args.end <= 0 {
		fmt.Fprintf(os.Stderr, "[Error: start || end must > 0]\n")
		os.Exit(1)
	} 
	if args.start > args.end {
		fmt.Fprintf(os.Stderr, "[Error: be sure end > start]\n")
		os.Exit(1)
	} 
	if args.pagetype == "f" && args.pagelen != 72 {
		fmt.Fprintf(os.Stderr, "[Error: if you choose -f and you can't use -l[number]]\n")
		os.Exit(1)
	} 
	//output the args 
	fmt.Println("start:", args.start);
	fmt.Println("end:", args.end);
	fmt.Println("pagelen:", args.pagelen);
	fmt.Println("page_type:", args.pagetype);
	fmt.Println("inputfile:", args.inputfile);
	fmt.Println("destfile:", args.dest);
	fmt.Println("other_ele:", args.other_ele);
}

func run(args *selgpNode) {
	var cmd_in io.WriteCloser
	if args.inputfile != "" { 
		content, err := os.Open(args.inputfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		linecnt := 1
		pagecnt := 1
		content_buf := bufio.NewReader(content)
		for {
			line, _, err := content_buf.ReadLine()
			if err != io.EOF && err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if err == io.EOF {
				break
			}
			// output to destfilename
			if pagecnt >= args.start && pagecnt <= args.end {
				if args.dest == "" {
					fmt.Println(string(line))
				} else {
					fmt.Fprintln(cmd_in, string(line))
				}
			}
			linecnt++
			if args.pagetype == "l" && linecnt > args.pagelen {
				linecnt = 1
				pagecnt++
			} else {
				if string(line) == "\f" {
					pagecnt++
				}
			}
		}
	} else { //input from the console
		_line := 1
		_page := 1
		for _,t := range args.other_ele {
			if _line > args.pagelen{
				_line = 1
				_page ++
			}
			if _page >= args.start && _page <= args.end {
				if args.dest == "" {
					fmt.Println(t)
				} else {
					fmt.Fprintln(cmd_in, t)
				}
			}
			_line ++;
		}
	}
}
