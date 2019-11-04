# Example WORKSPACE file for use with occt. The generator by default adds
# dependencies to @freetype2 and @tcl. You will likely need to change the paths
# for each of these repositories based on where your header files are.

new_local_repository(
    name = "freetype2",
    build_file_content = """
package(default_visibility=["//visibility:public"])
cc_library(
    name="freetype2",
    hdrs=glob([
        "*.h",
        "config/*.h",
        "freetype/*.h",
        "freetype/config/*.h",
    ]),
    includes=[".", "freetype"],
)
    """,
    path = "/usr/include/freetype2",
)

new_local_repository(
    name = "tcl",
    build_file_content = """
package(default_visibility=["//visibility:public"])
cc_library(
    name="tcl",
    hdrs=glob([
        "tcl*.h",
        "tk*.h",
    ]),
    includes = ["."],
    linkopts = ["-ltcl8.6"],
)
    """,
    path = "/usr/include/tcl8.6",
)
