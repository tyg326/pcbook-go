1.proto文件说明
    package 表示proto文件的包名

    option go_package = "path;name";
        path表示生成的go文件的存放地址，会自动生成目录的； 备注：该路径会在--go_out参数指定的路径下生成。推荐通过path指定
        name表示生成的go文件所属的包名
        option go_package = "./;teachschool_pcbook"; --go_out=./pb
        option go_package = "./pb;teachschool_pcbook"; --go_out=.

2.proto文件中import报红
    import报红是因为vscode-proto3扩展在运行protoc进行代码分析时使用我们当前的工作文件夹作为proto_path,所以无法在pcbook文件夹中找到memory_messagel.proto文件进行导入