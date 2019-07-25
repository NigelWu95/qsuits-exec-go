# qsuits-exec-go  
qsuits-java 代理执行工具 by golang  

Usage:  
      this tool is a agent program for qsuits, your local environment need java8 or above. In default mode, this tool will use latest java qsuits to exec, you only need use qsuits-java's parameters to run. If you use local mode it mean you dont want to update latest qsuits automatically.  
Options:  
        -Local          use current default qsuits version to exec.  
        --help/-h/help  print usage.  
Commands:  
         help           print usage.  
         versions       list all qsuits versions from local.  
         clear          remove all old qsuits versions from local.  
         current        query local default qsuits version.  
         chgver <no.>   set local default qsuits version.  
Usage of qsuits-java:  https://github.com/NigelWu95/qiniu-suits-java  
