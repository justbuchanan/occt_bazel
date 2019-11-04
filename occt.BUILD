package(default_visibility = ["//visibility:public"])

cc_library(
    name = "AIS",
    srcs = glob([
        "AIS/*.cxx",
    ]),
    hdrs = glob([
        "AIS/*.gxx",
        "AIS/*.pxx",
        "AIS/*.hxx",
        "AIS/*.lxx",
        "AIS/*.h",
    ]),
    includes = [
        "AIS/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":DsgPrs",
        ":ElCLib",
        ":ElSLib",
        ":Font",
        ":GC",
        ":GccEnt",
        ":GeomAbs",
        ":GeomLib",
        ":GeomProjLib",
        ":Image",
        ":IntAna2d",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":ProjLib",
        ":Quantity",
        ":StdPrs",
        ":StdSelect",
        ":TColgp",
        ":TShort",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Units",
        ":gp",
    ],
)

cc_library(
    name = "APIHeaderSection",
    srcs = glob([
        "APIHeaderSection/*.cxx",
    ]),
    hdrs = glob([
        "APIHeaderSection/*.gxx",
        "APIHeaderSection/*.pxx",
        "APIHeaderSection/*.hxx",
        "APIHeaderSection/*.lxx",
        "APIHeaderSection/*.h",
    ]),
    includes = [
        "APIHeaderSection/",
    ],
    deps = [
        ":HeaderSection",
        ":IFSelect",
        ":Interface",
        ":NCollection",
        ":StepData",
    ],
)

cc_library(
    name = "Adaptor2d",
    srcs = glob([
        "Adaptor2d/*.cxx",
        "Adaptor3d/*.cxx",
        "AdvApp2Var/*.cxx",
        "BndLib/*.cxx",
        "CPnts/*.cxx",
        "DrawTrSurf/*.cxx",
        "Extrema/*.cxx",
        "GCPnts/*.cxx",
        "Geom/*.cxx",
        "Geom2d/*.cxx",
        "Geom2dAdaptor/*.cxx",
        "Geom2dConvert/*.cxx",
        "Geom2dEvaluator/*.cxx",
        "Geom2dLProp/*.cxx",
        "GeomAdaptor/*.cxx",
        "GeomConvert/*.cxx",
        "GeomEvaluator/*.cxx",
        "GeomLProp/*.cxx",
        "GeomTools/*.cxx",
        "Hermit/*.cxx",
        "TColGeom/*.cxx",
        "TColGeom2d/*.cxx",
        "gce/*.cxx",
    ]),
    hdrs = glob([
        "Adaptor2d/*.gxx",
        "Adaptor2d/*.pxx",
        "Adaptor2d/*.hxx",
        "Adaptor2d/*.lxx",
        "Adaptor2d/*.h",
        "Adaptor3d/*.gxx",
        "Adaptor3d/*.pxx",
        "Adaptor3d/*.hxx",
        "Adaptor3d/*.lxx",
        "Adaptor3d/*.h",
        "AdvApp2Var/*.gxx",
        "AdvApp2Var/*.pxx",
        "AdvApp2Var/*.hxx",
        "AdvApp2Var/*.lxx",
        "AdvApp2Var/*.h",
        "BndLib/*.gxx",
        "BndLib/*.pxx",
        "BndLib/*.hxx",
        "BndLib/*.lxx",
        "BndLib/*.h",
        "CPnts/*.gxx",
        "CPnts/*.pxx",
        "CPnts/*.hxx",
        "CPnts/*.lxx",
        "CPnts/*.h",
        "DrawTrSurf/*.gxx",
        "DrawTrSurf/*.pxx",
        "DrawTrSurf/*.hxx",
        "DrawTrSurf/*.lxx",
        "DrawTrSurf/*.h",
        "Extrema/*.gxx",
        "Extrema/*.pxx",
        "Extrema/*.hxx",
        "Extrema/*.lxx",
        "Extrema/*.h",
        "GCPnts/*.gxx",
        "GCPnts/*.pxx",
        "GCPnts/*.hxx",
        "GCPnts/*.lxx",
        "GCPnts/*.h",
        "Geom/*.gxx",
        "Geom/*.pxx",
        "Geom/*.hxx",
        "Geom/*.lxx",
        "Geom/*.h",
        "Geom2d/*.gxx",
        "Geom2d/*.pxx",
        "Geom2d/*.hxx",
        "Geom2d/*.lxx",
        "Geom2d/*.h",
        "Geom2dAdaptor/*.gxx",
        "Geom2dAdaptor/*.pxx",
        "Geom2dAdaptor/*.hxx",
        "Geom2dAdaptor/*.lxx",
        "Geom2dAdaptor/*.h",
        "Geom2dConvert/*.gxx",
        "Geom2dConvert/*.pxx",
        "Geom2dConvert/*.hxx",
        "Geom2dConvert/*.lxx",
        "Geom2dConvert/*.h",
        "Geom2dEvaluator/*.gxx",
        "Geom2dEvaluator/*.pxx",
        "Geom2dEvaluator/*.hxx",
        "Geom2dEvaluator/*.lxx",
        "Geom2dEvaluator/*.h",
        "Geom2dLProp/*.gxx",
        "Geom2dLProp/*.pxx",
        "Geom2dLProp/*.hxx",
        "Geom2dLProp/*.lxx",
        "Geom2dLProp/*.h",
        "GeomAdaptor/*.gxx",
        "GeomAdaptor/*.pxx",
        "GeomAdaptor/*.hxx",
        "GeomAdaptor/*.lxx",
        "GeomAdaptor/*.h",
        "GeomConvert/*.gxx",
        "GeomConvert/*.pxx",
        "GeomConvert/*.hxx",
        "GeomConvert/*.lxx",
        "GeomConvert/*.h",
        "GeomEvaluator/*.gxx",
        "GeomEvaluator/*.pxx",
        "GeomEvaluator/*.hxx",
        "GeomEvaluator/*.lxx",
        "GeomEvaluator/*.h",
        "GeomLProp/*.gxx",
        "GeomLProp/*.pxx",
        "GeomLProp/*.hxx",
        "GeomLProp/*.lxx",
        "GeomLProp/*.h",
        "GeomTools/*.gxx",
        "GeomTools/*.pxx",
        "GeomTools/*.hxx",
        "GeomTools/*.lxx",
        "GeomTools/*.h",
        "Hermit/*.gxx",
        "Hermit/*.pxx",
        "Hermit/*.hxx",
        "Hermit/*.lxx",
        "Hermit/*.h",
        "TColGeom/*.gxx",
        "TColGeom/*.pxx",
        "TColGeom/*.hxx",
        "TColGeom/*.lxx",
        "TColGeom/*.h",
        "TColGeom2d/*.gxx",
        "TColGeom2d/*.pxx",
        "TColGeom2d/*.hxx",
        "TColGeom2d/*.lxx",
        "TColGeom2d/*.h",
        "gce/*.gxx",
        "gce/*.pxx",
        "gce/*.hxx",
        "gce/*.lxx",
        "gce/*.h",
    ]),
    includes = [
        "Adaptor2d/",
        "Adaptor3d/",
        "AdvApp2Var/",
        "BndLib/",
        "CPnts/",
        "DrawTrSurf/",
        "Extrema/",
        "GCPnts/",
        "Geom/",
        "Geom2d/",
        "Geom2dAdaptor/",
        "Geom2dConvert/",
        "Geom2dEvaluator/",
        "Geom2dLProp/",
        "GeomAdaptor/",
        "GeomConvert/",
        "GeomEvaluator/",
        "GeomLProp/",
        "GeomTools/",
        "Hermit/",
        "TColGeom/",
        "TColGeom2d/",
        "gce/",
    ],
    deps = [
        ":AdvApprox",
        ":BSplCLib",
        ":BSplSLib",
        ":Bnd",
        ":CSLib",
        ":Convert",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":IntAna",
        ":IntAna2d",
        ":LProp",
        ":Message",
        ":NCollection",
        ":PLib",
        ":Poly",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AdvApprox",
    srcs = glob([
        "AdvApprox/*.cxx",
    ]),
    hdrs = glob([
        "AdvApprox/*.gxx",
        "AdvApprox/*.pxx",
        "AdvApprox/*.hxx",
        "AdvApprox/*.lxx",
        "AdvApprox/*.h",
    ]),
    includes = [
        "AdvApprox/",
    ],
    deps = [
        ":BSplCLib",
        ":Convert",
        ":GeomAbs",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AppBlend",
    srcs = glob([
        "AppBlend/*.cxx",
    ]),
    hdrs = glob([
        "AppBlend/*.gxx",
        "AppBlend/*.pxx",
        "AppBlend/*.hxx",
        "AppBlend/*.lxx",
        "AppBlend/*.h",
    ]),
    includes = [
        "AppBlend/",
    ],
    deps = [
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":BSplCLib",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AppCont",
    srcs = glob([
        "AppCont/*.cxx",
    ]),
    hdrs = glob([
        "AppCont/*.gxx",
        "AppCont/*.pxx",
        "AppCont/*.hxx",
        "AppCont/*.lxx",
        "AppCont/*.h",
    ]),
    includes = [
        "AppCont/",
    ],
    deps = [
        ":AppParCurves",
        ":NCollection",
        ":PLib",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AppDef",
    srcs = glob([
        "AppDef/*.cxx",
    ]),
    hdrs = glob([
        "AppDef/*.gxx",
        "AppDef/*.pxx",
        "AppDef/*.hxx",
        "AppDef/*.lxx",
        "AppDef/*.h",
    ]),
    includes = [
        "AppDef/",
    ],
    deps = [
        ":AppParCurves",
        ":Approx",
        ":Convert",
        ":FEmTool",
        ":GeomAbs",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AppParCurves",
    srcs = glob([
        "AppParCurves/*.cxx",
    ]),
    hdrs = glob([
        "AppParCurves/*.gxx",
        "AppParCurves/*.pxx",
        "AppParCurves/*.hxx",
        "AppParCurves/*.lxx",
        "AppParCurves/*.h",
    ]),
    includes = [
        "AppParCurves/",
    ],
    deps = [
        ":BSplCLib",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "AppStd",
    srcs = glob([
        "AppStd/*.cxx",
    ]),
    hdrs = glob([
        "AppStd/*.gxx",
        "AppStd/*.pxx",
        "AppStd/*.hxx",
        "AppStd/*.lxx",
        "AppStd/*.h",
    ]),
    includes = [
        "AppStd/",
    ],
    deps = [
        ":NCollection",
        ":TDocStd",
    ],
)

cc_library(
    name = "AppStdL",
    srcs = glob([
        "AppStdL/*.cxx",
    ]),
    hdrs = glob([
        "AppStdL/*.gxx",
        "AppStdL/*.pxx",
        "AppStdL/*.hxx",
        "AppStdL/*.lxx",
        "AppStdL/*.h",
    ]),
    includes = [
        "AppStdL/",
    ],
    deps = [
        ":NCollection",
        ":TDocStd",
    ],
)

cc_library(
    name = "Approx",
    srcs = glob([
        "Approx/*.cxx",
    ]),
    hdrs = glob([
        "Approx/*.gxx",
        "Approx/*.pxx",
        "Approx/*.hxx",
        "Approx/*.lxx",
        "Approx/*.h",
    ]),
    includes = [
        "Approx/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":AppCont",
        ":AppParCurves",
        ":BSplCLib",
        ":Convert",
        ":Draw",
        ":GeomAbs",
        ":GeomLib",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "ApproxInt",
    srcs = glob([
        "ApproxInt/*.cxx",
        "BRepAdaptor/*.cxx",
        "BRepApprox/*.cxx",
        "BRepBndLib/*.cxx",
        "BRepBuilderAPI/*.cxx",
        "BRepCheck/*.cxx",
        "BRepClass/*.cxx",
        "BRepClass3d/*.cxx",
        "BRepExtrema/*.cxx",
        "BRepGProp/*.cxx",
        "BRepLib/*.cxx",
        "BRepMesh/*.cxx",
        "BRepTools/*.cxx",
        "BRepTopAdaptor/*.cxx",
        "BinTools/*.cxx",
        "Contap/*.cxx",
        "DBRep/*.cxx",
        "GeomAPI/*.cxx",
        "GeomInt/*.cxx",
        "HLRBRep/*.cxx",
        "HLRTopoBRep/*.cxx",
        "IntCurvesFace/*.cxx",
        "IntPatch/*.cxx",
        "IntStart/*.cxx",
    ]),
    hdrs = glob([
        "ApproxInt/*.gxx",
        "ApproxInt/*.pxx",
        "ApproxInt/*.hxx",
        "ApproxInt/*.lxx",
        "ApproxInt/*.h",
        "BRepAdaptor/*.gxx",
        "BRepAdaptor/*.pxx",
        "BRepAdaptor/*.hxx",
        "BRepAdaptor/*.lxx",
        "BRepAdaptor/*.h",
        "BRepApprox/*.gxx",
        "BRepApprox/*.pxx",
        "BRepApprox/*.hxx",
        "BRepApprox/*.lxx",
        "BRepApprox/*.h",
        "BRepBndLib/*.gxx",
        "BRepBndLib/*.pxx",
        "BRepBndLib/*.hxx",
        "BRepBndLib/*.lxx",
        "BRepBndLib/*.h",
        "BRepBuilderAPI/*.gxx",
        "BRepBuilderAPI/*.pxx",
        "BRepBuilderAPI/*.hxx",
        "BRepBuilderAPI/*.lxx",
        "BRepBuilderAPI/*.h",
        "BRepCheck/*.gxx",
        "BRepCheck/*.pxx",
        "BRepCheck/*.hxx",
        "BRepCheck/*.lxx",
        "BRepCheck/*.h",
        "BRepClass/*.gxx",
        "BRepClass/*.pxx",
        "BRepClass/*.hxx",
        "BRepClass/*.lxx",
        "BRepClass/*.h",
        "BRepClass3d/*.gxx",
        "BRepClass3d/*.pxx",
        "BRepClass3d/*.hxx",
        "BRepClass3d/*.lxx",
        "BRepClass3d/*.h",
        "BRepExtrema/*.gxx",
        "BRepExtrema/*.pxx",
        "BRepExtrema/*.hxx",
        "BRepExtrema/*.lxx",
        "BRepExtrema/*.h",
        "BRepGProp/*.gxx",
        "BRepGProp/*.pxx",
        "BRepGProp/*.hxx",
        "BRepGProp/*.lxx",
        "BRepGProp/*.h",
        "BRepLib/*.gxx",
        "BRepLib/*.pxx",
        "BRepLib/*.hxx",
        "BRepLib/*.lxx",
        "BRepLib/*.h",
        "BRepMesh/*.gxx",
        "BRepMesh/*.pxx",
        "BRepMesh/*.hxx",
        "BRepMesh/*.lxx",
        "BRepMesh/*.h",
        "BRepTools/*.gxx",
        "BRepTools/*.pxx",
        "BRepTools/*.hxx",
        "BRepTools/*.lxx",
        "BRepTools/*.h",
        "BRepTopAdaptor/*.gxx",
        "BRepTopAdaptor/*.pxx",
        "BRepTopAdaptor/*.hxx",
        "BRepTopAdaptor/*.lxx",
        "BRepTopAdaptor/*.h",
        "BinTools/*.gxx",
        "BinTools/*.pxx",
        "BinTools/*.hxx",
        "BinTools/*.lxx",
        "BinTools/*.h",
        "Contap/*.gxx",
        "Contap/*.pxx",
        "Contap/*.hxx",
        "Contap/*.lxx",
        "Contap/*.h",
        "DBRep/*.gxx",
        "DBRep/*.pxx",
        "DBRep/*.hxx",
        "DBRep/*.lxx",
        "DBRep/*.h",
        "GeomAPI/*.gxx",
        "GeomAPI/*.pxx",
        "GeomAPI/*.hxx",
        "GeomAPI/*.lxx",
        "GeomAPI/*.h",
        "GeomInt/*.gxx",
        "GeomInt/*.pxx",
        "GeomInt/*.hxx",
        "GeomInt/*.lxx",
        "GeomInt/*.h",
        "HLRBRep/*.gxx",
        "HLRBRep/*.pxx",
        "HLRBRep/*.hxx",
        "HLRBRep/*.lxx",
        "HLRBRep/*.h",
        "HLRTopoBRep/*.gxx",
        "HLRTopoBRep/*.pxx",
        "HLRTopoBRep/*.hxx",
        "HLRTopoBRep/*.lxx",
        "HLRTopoBRep/*.h",
        "IntCurvesFace/*.gxx",
        "IntCurvesFace/*.pxx",
        "IntCurvesFace/*.hxx",
        "IntCurvesFace/*.lxx",
        "IntCurvesFace/*.h",
        "IntPatch/*.gxx",
        "IntPatch/*.pxx",
        "IntPatch/*.hxx",
        "IntPatch/*.lxx",
        "IntPatch/*.h",
        "IntStart/*.gxx",
        "IntStart/*.pxx",
        "IntStart/*.hxx",
        "IntStart/*.lxx",
        "IntStart/*.h",
    ]),
    includes = [
        "ApproxInt/",
        "BRepAdaptor/",
        "BRepApprox/",
        "BRepBndLib/",
        "BRepBuilderAPI/",
        "BRepCheck/",
        "BRepClass/",
        "BRepClass3d/",
        "BRepExtrema/",
        "BRepGProp/",
        "BRepLib/",
        "BRepMesh/",
        "BRepTools/",
        "BRepTopAdaptor/",
        "BinTools/",
        "Contap/",
        "DBRep/",
        "GeomAPI/",
        "GeomInt/",
        "HLRBRep/",
        "HLRTopoBRep/",
        "IntCurvesFace/",
        "IntPatch/",
        "IntStart/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":BRep",
        ":BSplCLib",
        ":BVH",
        ":Bnd",
        ":CSLib",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":FSD",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dHatch",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":GeomLib",
        ":GeomProjLib",
        ":HLRAlgo",
        ":HatchGen",
        ":IntAna",
        ":IntCurve",
        ":IntCurveSurface",
        ":IntImp",
        ":IntImpParGen",
        ":IntPolyh",
        ":IntRes2d",
        ":IntSurf",
        ":IntWalk",
        ":Intf",
        ":LProp",
        ":LocalAnalysis",
        ":Message",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":Plugin",
        ":Poly",
        ":ProjLib",
        ":StdFail",
        ":Storage",
        ":TColgp",
        ":TShort",
        ":TopAbs",
        ":TopClass",
        ":TopCnx",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopTrans",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Aspect",
    srcs = glob([
        "Aspect/*.cxx",
    ]),
    hdrs = glob([
        "Aspect/*.gxx",
        "Aspect/*.pxx",
        "Aspect/*.hxx",
        "Aspect/*.lxx",
        "Aspect/*.h",
    ]),
    includes = [
        "Aspect/",
    ],
    deps = [
        ":InterfaceGraphic",
        ":NCollection",
        ":OSD",
        ":Quantity",
    ],
)

cc_library(
    name = "BOPAlgo",
    srcs = glob([
        "BOPAlgo/*.cxx",
        "BOPDS/*.cxx",
        "BOPTools/*.cxx",
    ]),
    hdrs = glob([
        "BOPAlgo/*.gxx",
        "BOPAlgo/*.pxx",
        "BOPAlgo/*.hxx",
        "BOPAlgo/*.lxx",
        "BOPAlgo/*.h",
        "BOPDS/*.gxx",
        "BOPDS/*.pxx",
        "BOPDS/*.hxx",
        "BOPDS/*.lxx",
        "BOPDS/*.h",
        "BOPTools/*.gxx",
        "BOPTools/*.pxx",
        "BOPTools/*.hxx",
        "BOPTools/*.lxx",
        "BOPTools/*.h",
    ]),
    includes = [
        "BOPAlgo/",
        "BOPDS/",
        "BOPTools/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepPrimAPI",
        ":Bnd",
        ":ElCLib",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dHatch",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomLib",
        ":GeomProjLib",
        ":HatchGen",
        ":IntRes2d",
        ":IntSurf",
        ":IntTools",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":ProjLib",
        ":ShapeUpgrade",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BOPTest",
    srcs = glob([
        "BOPTest/*.cxx",
    ]),
    hdrs = glob([
        "BOPTest/*.gxx",
        "BOPTest/*.pxx",
        "BOPTest/*.hxx",
        "BOPTest/*.lxx",
        "BOPTest/*.h",
    ]),
    includes = [
        "BOPTest/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgoAPI",
        ":BRepTest",
        ":Draw",
        ":GeometryTest",
        ":GeomliteTest",
        ":HLRTest",
        ":IntTools",
        ":MeshTest",
        ":Message",
        ":NCollection",
        ":OSD",
        ":SWDRAW",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRep",
    srcs = glob([
        "BRep/*.cxx",
    ]),
    hdrs = glob([
        "BRep/*.gxx",
        "BRep/*.pxx",
        "BRep/*.hxx",
        "BRep/*.lxx",
        "BRep/*.h",
    ]),
    includes = [
        "BRep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElSLib",
        ":GeomAbs",
        ":GeomProjLib",
        ":NCollection",
        ":Poly",
        ":ProjLib",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepAlgo",
    srcs = glob([
        "BRepAlgo/*.cxx",
    ]),
    hdrs = glob([
        "BRepAlgo/*.gxx",
        "BRepAlgo/*.pxx",
        "BRepAlgo/*.hxx",
        "BRepAlgo/*.lxx",
        "BRepAlgo/*.h",
    ]),
    includes = [
        "BRepAlgo/",
    ],
    deps = [
        ":Adaptor2d",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgoAPI",
        ":Bnd",
        ":ElCLib",
        ":GeomAbs",
        ":GeomProjLib",
        ":NCollection",
        ":OSD",
        ":ProjLib",
        ":ShapeAnalysis",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopOpeBRep",
        ":TopOpeBRepBuild",
        ":TopOpeBRepDS",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepAlgoAPI",
    srcs = glob([
        "BRepAlgoAPI/*.cxx",
    ]),
    hdrs = glob([
        "BRepAlgoAPI/*.gxx",
        "BRepAlgoAPI/*.pxx",
        "BRepAlgoAPI/*.hxx",
        "BRepAlgoAPI/*.lxx",
        "BRepAlgoAPI/*.h",
    ]),
    includes = [
        "BRepAlgoAPI/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BOPAlgo",
        ":NCollection",
        ":OSD",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepBlend",
    srcs = glob([
        "BRepBlend/*.cxx",
    ]),
    hdrs = glob([
        "BRepBlend/*.gxx",
        "BRepBlend/*.pxx",
        "BRepBlend/*.hxx",
        "BRepBlend/*.lxx",
        "BRepBlend/*.h",
    ]),
    includes = [
        "BRepBlend/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppBlend",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":Blend",
        ":BlendFunc",
        ":ChFiDS",
        ":Convert",
        ":ElCLib",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":IntRes2d",
        ":IntSurf",
        ":Law",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "BRepFeat",
    srcs = glob([
        "BRepFeat/*.cxx",
    ]),
    hdrs = glob([
        "BRepFeat/*.gxx",
        "BRepFeat/*.pxx",
        "BRepFeat/*.hxx",
        "BRepFeat/*.lxx",
        "BRepFeat/*.h",
    ]),
    includes = [
        "BRepFeat/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":BRepIntCurveSurface",
        ":BRepPrim",
        ":BRepPrimAPI",
        ":BRepSweep",
        ":Bnd",
        ":CSLib",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomLib",
        ":GeomProjLib",
        ":IntRes2d",
        ":IntTools",
        ":LocOpe",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepFill",
    srcs = glob([
        "BRepFill/*.cxx",
    ]),
    hdrs = glob([
        "BRepFill/*.gxx",
        "BRepFill/*.pxx",
        "BRepFill/*.hxx",
        "BRepFill/*.lxx",
        "BRepFill/*.h",
    ]),
    includes = [
        "BRepFill/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppCont",
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":ApproxInt",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":BRepIntCurveSurface",
        ":BRepLProp",
        ":BRepMAT2d",
        ":BRepSweep",
        ":BSplCLib",
        ":Bisector",
        ":Bnd",
        ":Draw",
        ":ElCLib",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":GeomLib",
        ":GeomPlate",
        ":GeomProjLib",
        ":IntCurveSurface",
        ":IntRes2d",
        ":IntTools",
        ":Law",
        ":MAT",
        ":MAT2d",
        ":NCollection",
        ":PLib",
        ":ProjLib",
        ":ShapeUpgrade",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepFilletAPI",
    srcs = glob([
        "BRepFilletAPI/*.cxx",
    ]),
    hdrs = glob([
        "BRepFilletAPI/*.gxx",
        "BRepFilletAPI/*.pxx",
        "BRepFilletAPI/*.hxx",
        "BRepFilletAPI/*.lxx",
        "BRepFilletAPI/*.h",
    ]),
    includes = [
        "BRepFilletAPI/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":ChFi2d",
        ":ChFi3d",
        ":ChFiDS",
        ":GeomAbs",
        ":Law",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":TopExp",
        ":TopOpeBRepBuild",
        ":TopOpeBRepDS",
        ":TopTools",
        ":TopoDS",
    ],
)

cc_library(
    name = "BRepIntCurveSurface",
    srcs = glob([
        "BRepIntCurveSurface/*.cxx",
    ]),
    hdrs = glob([
        "BRepIntCurveSurface/*.gxx",
        "BRepIntCurveSurface/*.pxx",
        "BRepIntCurveSurface/*.hxx",
        "BRepIntCurveSurface/*.lxx",
        "BRepIntCurveSurface/*.h",
    ]),
    includes = [
        "BRepIntCurveSurface/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Bnd",
        ":IntCurveSurface",
        ":NCollection",
        ":StdFail",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepLProp",
    srcs = glob([
        "BRepLProp/*.cxx",
    ]),
    hdrs = glob([
        "BRepLProp/*.gxx",
        "BRepLProp/*.pxx",
        "BRepLProp/*.hxx",
        "BRepLProp/*.lxx",
        "BRepLProp/*.h",
    ]),
    includes = [
        "BRepLProp/",
    ],
    deps = [
        ":ApproxInt",
        ":GeomAbs",
        ":LProp",
        ":NCollection",
        ":TopAbs",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepMAT2d",
    srcs = glob([
        "BRepMAT2d/*.cxx",
    ]),
    hdrs = glob([
        "BRepMAT2d/*.gxx",
        "BRepMAT2d/*.pxx",
        "BRepMAT2d/*.hxx",
        "BRepMAT2d/*.lxx",
        "BRepMAT2d/*.h",
    ]),
    includes = [
        "BRepMAT2d/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Bisector",
        ":GCE2d",
        ":GeomAbs",
        ":MAT",
        ":MAT2d",
        ":NCollection",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepOffset",
    srcs = glob([
        "BRepOffset/*.cxx",
    ]),
    hdrs = glob([
        "BRepOffset/*.gxx",
        "BRepOffset/*.pxx",
        "BRepOffset/*.hxx",
        "BRepOffset/*.lxx",
        "BRepOffset/*.h",
    ]),
    includes = [
        "BRepOffset/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":Bnd",
        ":ElCLib",
        ":ElSLib",
        ":GC",
        ":GCE2d",
        ":GProp",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":GeomLib",
        ":GeomProjLib",
        ":IntRes2d",
        ":IntTools",
        ":NCollection",
        ":OSD",
        ":ProjLib",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepOffsetAPI",
    srcs = glob([
        "BRepOffsetAPI/*.cxx",
    ]),
    hdrs = glob([
        "BRepOffsetAPI/*.gxx",
        "BRepOffsetAPI/*.pxx",
        "BRepOffsetAPI/*.hxx",
        "BRepOffsetAPI/*.lxx",
        "BRepOffsetAPI/*.h",
    ]),
    includes = [
        "BRepOffsetAPI/",
    ],
    deps = [
        ":Adaptor2d",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgo",
        ":BRepFill",
        ":BRepOffset",
        ":BRepPrimAPI",
        ":BSplCLib",
        ":Draft",
        ":GC",
        ":GCE2d",
        ":GProp",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":GeomLib",
        ":IntRes2d",
        ":Law",
        ":NCollection",
        ":ShapeUpgrade",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepPrim",
    srcs = glob([
        "BRepPrim/*.cxx",
    ]),
    hdrs = glob([
        "BRepPrim/*.gxx",
        "BRepPrim/*.pxx",
        "BRepPrim/*.hxx",
        "BRepPrim/*.lxx",
        "BRepPrim/*.h",
    ]),
    includes = [
        "BRepPrim/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":ElSLib",
        ":NCollection",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepPrimAPI",
    srcs = glob([
        "BRepPrimAPI/*.cxx",
    ]),
    hdrs = glob([
        "BRepPrimAPI/*.gxx",
        "BRepPrimAPI/*.pxx",
        "BRepPrimAPI/*.hxx",
        "BRepPrimAPI/*.lxx",
        "BRepPrimAPI/*.h",
    ]),
    includes = [
        "BRepPrimAPI/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepLProp",
        ":BRepPrim",
        ":BRepSweep",
        ":GeomProjLib",
        ":NCollection",
        ":StdFail",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepProj",
    srcs = glob([
        "BRepProj/*.cxx",
    ]),
    hdrs = glob([
        "BRepProj/*.gxx",
        "BRepProj/*.pxx",
        "BRepProj/*.hxx",
        "BRepProj/*.lxx",
        "BRepProj/*.h",
    ]),
    includes = [
        "BRepProj/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":BRepAlgoAPI",
        ":BRepFill",
        ":BRepSweep",
        ":Bnd",
        ":NCollection",
        ":ShapeAnalysis",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepSweep",
    srcs = glob([
        "BRepSweep/*.cxx",
    ]),
    hdrs = glob([
        "BRepSweep/*.gxx",
        "BRepSweep/*.pxx",
        "BRepSweep/*.hxx",
        "BRepSweep/*.lxx",
        "BRepSweep/*.h",
    ]),
    includes = [
        "BRepSweep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepLProp",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":NCollection",
        ":Sweep",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepTest",
    srcs = glob([
        "BRepTest/*.cxx",
    ]),
    hdrs = glob([
        "BRepTest/*.gxx",
        "BRepTest/*.pxx",
        "BRepTest/*.hxx",
        "BRepTest/*.lxx",
        "BRepTest/*.h",
    ]),
    includes = [
        "BRepTest/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":ApproxInt",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":BRepFeat",
        ":BRepFill",
        ":BRepFilletAPI",
        ":BRepIntCurveSurface",
        ":BRepMAT2d",
        ":BRepOffset",
        ":BRepOffsetAPI",
        ":BRepPrimAPI",
        ":BRepProj",
        ":BiTgte",
        ":Bisector",
        ":Bnd",
        ":ChFi2d",
        ":ChFi3d",
        ":Draw",
        ":ElCLib",
        ":FilletSurf",
        ":GProp",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomFill",
        ":GeomPlate",
        ":GeometryTest",
        ":Law",
        ":LocOpe",
        ":LocalAnalysis",
        ":MAT",
        ":NCollection",
        ":OSD",
        ":ProjLib",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopOpeBRep",
        ":TopOpeBRepBuild",
        ":TopOpeBRepDS",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BRepToIGES",
    srcs = glob([
        "BRepToIGES/*.cxx",
    ]),
    hdrs = glob([
        "BRepToIGES/*.gxx",
        "BRepToIGES/*.pxx",
        "BRepToIGES/*.hxx",
        "BRepToIGES/*.lxx",
        "BRepToIGES/*.h",
    ]),
    includes = [
        "BRepToIGES/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Geom2dToIGES",
        ":GeomToIGES",
        ":IGESBasic",
        ":Interface",
        ":Message",
        ":NCollection",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":gp",
    ],
)

cc_library(
    name = "BRepToIGESBRep",
    srcs = glob([
        "BRepToIGESBRep/*.cxx",
    ]),
    hdrs = glob([
        "BRepToIGESBRep/*.gxx",
        "BRepToIGESBRep/*.pxx",
        "BRepToIGESBRep/*.hxx",
        "BRepToIGESBRep/*.lxx",
        "BRepToIGESBRep/*.h",
    ]),
    includes = [
        "BRepToIGESBRep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepToIGES",
        ":Geom2dToIGES",
        ":GeomToIGES",
        ":IGESBasic",
        ":IGESSolid",
        ":Interface",
        ":Message",
        ":NCollection",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeCustom",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":gp",
    ],
)

cc_library(
    name = "BSplCLib",
    srcs = glob([
        "BSplCLib/*.cxx",
    ]),
    hdrs = glob([
        "BSplCLib/*.gxx",
        "BSplCLib/*.pxx",
        "BSplCLib/*.hxx",
        "BSplCLib/*.lxx",
        "BSplCLib/*.h",
    ]),
    includes = [
        "BSplCLib/",
    ],
    deps = [
        ":ElCLib",
        ":GeomAbs",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "BSplSLib",
    srcs = glob([
        "BSplSLib/*.cxx",
    ]),
    hdrs = glob([
        "BSplSLib/*.gxx",
        "BSplSLib/*.pxx",
        "BSplSLib/*.hxx",
        "BSplSLib/*.lxx",
        "BSplSLib/*.h",
    ]),
    includes = [
        "BSplSLib/",
    ],
    deps = [
        ":BSplCLib",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "BVH",
    srcs = glob([
        "BVH/*.cxx",
    ]),
    hdrs = glob([
        "BVH/*.gxx",
        "BVH/*.pxx",
        "BVH/*.hxx",
        "BVH/*.lxx",
        "BVH/*.h",
    ]),
    includes = [
        "BVH/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "BiTgte",
    srcs = glob([
        "BiTgte/*.cxx",
    ]),
    hdrs = glob([
        "BiTgte/*.gxx",
        "BiTgte/*.pxx",
        "BiTgte/*.hxx",
        "BiTgte/*.lxx",
        "BiTgte/*.h",
    ]),
    includes = [
        "BiTgte/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppCont",
        ":AppParCurves",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgo",
        ":BRepFill",
        ":BRepOffset",
        ":BSplCLib",
        ":Bnd",
        ":ChFi3d",
        ":Convert",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":NCollection",
        ":OSD",
        ":StdFail",
        ":TColgp",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "BinDrivers",
    srcs = glob([
        "BinDrivers/*.cxx",
    ]),
    hdrs = glob([
        "BinDrivers/*.gxx",
        "BinDrivers/*.pxx",
        "BinDrivers/*.hxx",
        "BinDrivers/*.lxx",
        "BinDrivers/*.h",
    ]),
    includes = [
        "BinDrivers/",
    ],
    deps = [
        ":BinLDrivers",
        ":BinMDF",
        ":BinMDataStd",
        ":BinMDataXtd",
        ":BinMDocStd",
        ":BinMFunction",
        ":BinMNaming",
        ":Message",
        ":NCollection",
        ":Plugin",
        ":Storage",
        ":TDocStd",
        ":TNaming",
    ],
)

cc_library(
    name = "BinLDrivers",
    srcs = glob([
        "BinLDrivers/*.cxx",
    ]),
    hdrs = glob([
        "BinLDrivers/*.gxx",
        "BinLDrivers/*.pxx",
        "BinLDrivers/*.hxx",
        "BinLDrivers/*.lxx",
        "BinLDrivers/*.h",
    ]),
    includes = [
        "BinLDrivers/",
    ],
    deps = [
        ":BinMDF",
        ":BinMDataStd",
        ":BinMDocStd",
        ":BinMFunction",
        ":BinMNaming",
        ":BinObjMgt",
        ":CDM",
        ":FSD",
        ":Message",
        ":NCollection",
        ":OSD",
        ":PCDM",
        ":Plugin",
        ":Storage",
        ":TDF",
        ":TDocStd",
    ],
)

cc_library(
    name = "BinMDF",
    srcs = glob([
        "BinMDF/*.cxx",
    ]),
    hdrs = glob([
        "BinMDF/*.gxx",
        "BinMDF/*.pxx",
        "BinMDF/*.hxx",
        "BinMDF/*.lxx",
        "BinMDF/*.h",
    ]),
    includes = [
        "BinMDF/",
    ],
    deps = [
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
    ],
)

cc_library(
    name = "BinMDataStd",
    srcs = glob([
        "BinMDataStd/*.cxx",
    ]),
    hdrs = glob([
        "BinMDataStd/*.gxx",
        "BinMDataStd/*.pxx",
        "BinMDataStd/*.hxx",
        "BinMDataStd/*.lxx",
        "BinMDataStd/*.h",
    ]),
    includes = [
        "BinMDataStd/",
    ],
    deps = [
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TDataStd",
    ],
)

cc_library(
    name = "BinMDataXtd",
    srcs = glob([
        "BinMDataXtd/*.cxx",
    ]),
    hdrs = glob([
        "BinMDataXtd/*.gxx",
        "BinMDataXtd/*.pxx",
        "BinMDataXtd/*.hxx",
        "BinMDataXtd/*.lxx",
        "BinMDataXtd/*.h",
    ]),
    includes = [
        "BinMDataXtd/",
    ],
    deps = [
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TDataStd",
        ":TDataXtd",
        ":TNaming",
        ":gp",
    ],
)

cc_library(
    name = "BinMDocStd",
    srcs = glob([
        "BinMDocStd/*.cxx",
    ]),
    hdrs = glob([
        "BinMDocStd/*.gxx",
        "BinMDocStd/*.pxx",
        "BinMDocStd/*.hxx",
        "BinMDocStd/*.lxx",
        "BinMDocStd/*.h",
    ]),
    includes = [
        "BinMDocStd/",
    ],
    deps = [
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TDocStd",
    ],
)

cc_library(
    name = "BinMFunction",
    srcs = glob([
        "BinMFunction/*.cxx",
    ]),
    hdrs = glob([
        "BinMFunction/*.gxx",
        "BinMFunction/*.pxx",
        "BinMFunction/*.hxx",
        "BinMFunction/*.lxx",
        "BinMFunction/*.h",
    ]),
    includes = [
        "BinMFunction/",
    ],
    deps = [
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TFunction",
    ],
)

cc_library(
    name = "BinMNaming",
    srcs = glob([
        "BinMNaming/*.cxx",
    ]),
    hdrs = glob([
        "BinMNaming/*.gxx",
        "BinMNaming/*.pxx",
        "BinMNaming/*.hxx",
        "BinMNaming/*.lxx",
        "BinMNaming/*.h",
    ]),
    includes = [
        "BinMNaming/",
    ],
    deps = [
        ":ApproxInt",
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TNaming",
        ":TopAbs",
        ":TopoDS",
    ],
)

cc_library(
    name = "BinMXCAFDoc",
    srcs = glob([
        "BinMXCAFDoc/*.cxx",
    ]),
    hdrs = glob([
        "BinMXCAFDoc/*.gxx",
        "BinMXCAFDoc/*.pxx",
        "BinMXCAFDoc/*.hxx",
        "BinMXCAFDoc/*.lxx",
        "BinMXCAFDoc/*.h",
    ]),
    includes = [
        "BinMXCAFDoc/",
    ],
    deps = [
        ":ApproxInt",
        ":BinMDF",
        ":BinMDataStd",
        ":BinMNaming",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TNaming",
        ":TopLoc",
        ":XCAFDimTolObjects",
        ":gp",
    ],
)

cc_library(
    name = "BinObjMgt",
    srcs = glob([
        "BinObjMgt/*.cxx",
    ]),
    hdrs = glob([
        "BinObjMgt/*.gxx",
        "BinObjMgt/*.pxx",
        "BinObjMgt/*.hxx",
        "BinObjMgt/*.lxx",
        "BinObjMgt/*.h",
    ]),
    includes = [
        "BinObjMgt/",
    ],
    deps = [
        ":FSD",
        ":NCollection",
        ":Storage",
        ":TDF",
    ],
)

cc_library(
    name = "BinTObjDrivers",
    srcs = glob([
        "BinTObjDrivers/*.cxx",
    ]),
    hdrs = glob([
        "BinTObjDrivers/*.gxx",
        "BinTObjDrivers/*.pxx",
        "BinTObjDrivers/*.hxx",
        "BinTObjDrivers/*.lxx",
        "BinTObjDrivers/*.h",
    ]),
    includes = [
        "BinTObjDrivers/",
    ],
    deps = [
        ":BinLDrivers",
        ":BinMDF",
        ":BinObjMgt",
        ":Message",
        ":NCollection",
        ":Plugin",
        ":TDF",
        ":TDocStd",
        ":TObj",
    ],
)

cc_library(
    name = "BinXCAFDrivers",
    srcs = glob([
        "BinXCAFDrivers/*.cxx",
    ]),
    hdrs = glob([
        "BinXCAFDrivers/*.gxx",
        "BinXCAFDrivers/*.pxx",
        "BinXCAFDrivers/*.hxx",
        "BinXCAFDrivers/*.lxx",
        "BinXCAFDrivers/*.h",
    ]),
    includes = [
        "BinXCAFDrivers/",
    ],
    deps = [
        ":BinDrivers",
        ":BinMDF",
        ":BinMXCAFDoc",
        ":Message",
        ":NCollection",
        ":Plugin",
        ":TDocStd",
    ],
)

cc_library(
    name = "Bisector",
    srcs = glob([
        "Bisector/*.cxx",
    ]),
    hdrs = glob([
        "Bisector/*.gxx",
        "Bisector/*.pxx",
        "Bisector/*.hxx",
        "Bisector/*.lxx",
        "Bisector/*.h",
    ]),
    includes = [
        "Bisector/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":GCE2d",
        ":GccAna",
        ":GccEnt",
        ":GccInt",
        ":Geom2dAPI",
        ":Geom2dGcc",
        ":Geom2dInt",
        ":GeomAbs",
        ":IntAna2d",
        ":IntRes2d",
        ":NCollection",
        ":StdFail",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Blend",
    srcs = glob([
        "Blend/*.cxx",
    ]),
    hdrs = glob([
        "Blend/*.gxx",
        "Blend/*.pxx",
        "Blend/*.hxx",
        "Blend/*.lxx",
        "Blend/*.h",
    ]),
    includes = [
        "Blend/",
    ],
    deps = [
        ":Adaptor2d",
        ":CSLib",
        ":Draw",
        ":GeomAbs",
        ":IntSurf",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "BlendFunc",
    srcs = glob([
        "BlendFunc/*.cxx",
    ]),
    hdrs = glob([
        "BlendFunc/*.gxx",
        "BlendFunc/*.pxx",
        "BlendFunc/*.hxx",
        "BlendFunc/*.lxx",
        "BlendFunc/*.h",
    ]),
    includes = [
        "BlendFunc/",
    ],
    deps = [
        ":Adaptor2d",
        ":Blend",
        ":CSLib",
        ":Convert",
        ":ElCLib",
        ":GeomAbs",
        ":GeomFill",
        ":Law",
        ":NCollection",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Bnd",
    srcs = glob([
        "Bnd/*.cxx",
    ]),
    hdrs = glob([
        "Bnd/*.gxx",
        "Bnd/*.pxx",
        "Bnd/*.hxx",
        "Bnd/*.lxx",
        "Bnd/*.h",
    ]),
    includes = [
        "Bnd/",
    ],
    deps = [
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "CDF",
    srcs = glob([
        "CDF/*.cxx",
    ]),
    hdrs = glob([
        "CDF/*.gxx",
        "CDF/*.pxx",
        "CDF/*.hxx",
        "CDF/*.lxx",
        "CDF/*.h",
    ]),
    includes = [
        "CDF/",
    ],
    deps = [
        ":CDM",
        ":NCollection",
        ":OSD",
        ":PCDM",
        ":Plugin",
        ":UTL",
    ],
)

cc_library(
    name = "CDM",
    srcs = glob([
        "CDM/*.cxx",
    ]),
    hdrs = glob([
        "CDM/*.gxx",
        "CDM/*.pxx",
        "CDM/*.hxx",
        "CDM/*.lxx",
        "CDM/*.h",
    ]),
    includes = [
        "CDM/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":Resource",
        ":UTL",
    ],
)

cc_library(
    name = "CSLib",
    srcs = glob([
        "CSLib/*.cxx",
    ]),
    hdrs = glob([
        "CSLib/*.gxx",
        "CSLib/*.pxx",
        "CSLib/*.hxx",
        "CSLib/*.lxx",
        "CSLib/*.h",
    ]),
    includes = [
        "CSLib/",
    ],
    deps = [
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "ChFi2d",
    srcs = glob([
        "ChFi2d/*.cxx",
    ]),
    hdrs = glob([
        "ChFi2d/*.gxx",
        "ChFi2d/*.pxx",
        "ChFi2d/*.hxx",
        "ChFi2d/*.lxx",
        "ChFi2d/*.h",
    ]),
    includes = [
        "ChFi2d/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":ElSLib",
        ":GC",
        ":GccEnt",
        ":Geom2dAPI",
        ":Geom2dGcc",
        ":Geom2dInt",
        ":GeomProjLib",
        ":IntAna2d",
        ":IntRes2d",
        ":NCollection",
        ":ProjLib",
        ":ShapeAnalysis",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ChFi3d",
    srcs = glob([
        "ChFi3d/*.cxx",
    ]),
    hdrs = glob([
        "ChFi3d/*.gxx",
        "ChFi3d/*.pxx",
        "ChFi3d/*.hxx",
        "ChFi3d/*.lxx",
        "ChFi3d/*.h",
    ]),
    includes = [
        "ChFi3d/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppBlend",
        ":AppParCurves",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgo",
        ":BRepBlend",
        ":BRepLProp",
        ":BSplCLib",
        ":Blend",
        ":BlendFunc",
        ":Bnd",
        ":ChFiDS",
        ":ChFiKPart",
        ":Convert",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":FairCurve",
        ":GC",
        ":Geom2dAPI",
        ":Geom2dHatch",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomFill",
        ":GeomLib",
        ":GeomPlate",
        ":GeomProjLib",
        ":HatchGen",
        ":IntAna",
        ":IntAna2d",
        ":IntCurveSurface",
        ":IntRes2d",
        ":IntSurf",
        ":IntWalk",
        ":Law",
        ":LocalAnalysis",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":ProjLib",
        ":ShapeAnalysis",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopOpeBRepBuild",
        ":TopOpeBRepDS",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "ChFiDS",
    srcs = glob([
        "ChFiDS/*.cxx",
    ]),
    hdrs = glob([
        "ChFiDS/*.gxx",
        "ChFiDS/*.pxx",
        "ChFiDS/*.hxx",
        "ChFiDS/*.lxx",
        "ChFiDS/*.h",
    ]),
    includes = [
        "ChFiDS/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":GeomAbs",
        ":Law",
        ":NCollection",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ChFiKPart",
    srcs = glob([
        "ChFiKPart/*.cxx",
    ]),
    hdrs = glob([
        "ChFiKPart/*.gxx",
        "ChFiKPart/*.pxx",
        "ChFiKPart/*.hxx",
        "ChFiKPart/*.lxx",
        "ChFiKPart/*.h",
    ]),
    includes = [
        "ChFiKPart/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":ChFiDS",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dInt",
        ":GeomAbs",
        ":IntAna",
        ":IntRes2d",
        ":IntSurf",
        ":NCollection",
        ":ProjLib",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopOpeBRepDS",
        ":gp",
    ],
)

cc_library(
    name = "Cocoa",
    srcs = glob([
        "Cocoa/*.cxx",
    ]),
    hdrs = glob([
        "Cocoa/*.gxx",
        "Cocoa/*.pxx",
        "Cocoa/*.hxx",
        "Cocoa/*.lxx",
        "Cocoa/*.h",
    ]),
    includes = [
        "Cocoa/",
    ],
    deps = [
        ":Aspect",
        ":Image",
        ":NCollection",
        ":Quantity",
    ],
)

cc_library(
    name = "Convert",
    srcs = glob([
        "Convert/*.cxx",
    ]),
    hdrs = glob([
        "Convert/*.gxx",
        "Convert/*.pxx",
        "Convert/*.hxx",
        "Convert/*.lxx",
        "Convert/*.h",
    ]),
    includes = [
        "Convert/",
    ],
    deps = [
        ":BSplCLib",
        ":BSplSLib",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "D3DHost",
    srcs = glob([
        "D3DHost/*.cxx",
    ]),
    hdrs = glob([
        "D3DHost/*.gxx",
        "D3DHost/*.pxx",
        "D3DHost/*.hxx",
        "D3DHost/*.lxx",
        "D3DHost/*.h",
    ]),
    includes = [
        "D3DHost/",
    ],
    deps = [
        ":NCollection",
        ":OpenGl",
    ],
)

cc_library(
    name = "DDF",
    srcs = glob([
        "DDF/*.cxx",
    ]),
    hdrs = glob([
        "DDF/*.gxx",
        "DDF/*.pxx",
        "DDF/*.hxx",
        "DDF/*.lxx",
        "DDF/*.h",
    ]),
    includes = [
        "DDF/",
    ],
    deps = [
        ":Draw",
        ":NCollection",
        ":OSD",
        ":Storage",
        ":TDF",
        ":TDataStd",
    ],
)

cc_library(
    name = "DDataStd",
    srcs = glob([
        "DDataStd/*.cxx",
    ]),
    hdrs = glob([
        "DDataStd/*.gxx",
        "DDataStd/*.pxx",
        "DDataStd/*.hxx",
        "DDataStd/*.lxx",
        "DDataStd/*.h",
    ]),
    includes = [
        "DDataStd/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":DDF",
        ":Draw",
        ":DrawDim",
        ":ElSLib",
        ":NCollection",
        ":TDF",
        ":TDataStd",
        ":TDataXtd",
        ":TFunction",
        ":TNaming",
        ":TopAbs",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "DDocStd",
    srcs = glob([
        "DDocStd/*.cxx",
    ]),
    hdrs = glob([
        "DDocStd/*.gxx",
        "DDocStd/*.pxx",
        "DDocStd/*.hxx",
        "DDocStd/*.lxx",
        "DDocStd/*.h",
    ]),
    includes = [
        "DDocStd/",
    ],
    deps = [
        ":AIS",
        ":ApproxInt",
        ":BRep",
        ":BinDrivers",
        ":BinLDrivers",
        ":CDF",
        ":CDM",
        ":DDF",
        ":Draw",
        ":FSD",
        ":Font",
        ":NCollection",
        ":OSD",
        ":Plugin",
        ":ShapePersistent",
        ":StdDrivers",
        ":StdLDrivers",
        ":StdStorage",
        ":Storage",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TNaming",
        ":TPrsStd",
        ":TopTools",
        ":TopoDS",
        ":ViewerTest",
        ":XmlDrivers",
        ":XmlLDrivers",
    ],
)

cc_library(
    name = "DNaming",
    srcs = glob([
        "DNaming/*.cxx",
    ]),
    hdrs = glob([
        "DNaming/*.gxx",
        "DNaming/*.pxx",
        "DNaming/*.hxx",
        "DNaming/*.lxx",
        "DNaming/*.h",
    ]),
    includes = [
        "DNaming/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":BRepFilletAPI",
        ":BRepPrimAPI",
        ":DDF",
        ":DDocStd",
        ":Draw",
        ":GProp",
        ":GeomAbs",
        ":GeomLib",
        ":NCollection",
        ":TColgp",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TFunction",
        ":TNaming",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "DPrsStd",
    srcs = glob([
        "DPrsStd/*.cxx",
    ]),
    hdrs = glob([
        "DPrsStd/*.gxx",
        "DPrsStd/*.pxx",
        "DPrsStd/*.hxx",
        "DPrsStd/*.lxx",
        "DPrsStd/*.h",
    ]),
    includes = [
        "DPrsStd/",
    ],
    deps = [
        ":AIS",
        ":AppStd",
        ":DDF",
        ":DDataStd",
        ":DDocStd",
        ":DNaming",
        ":Draw",
        ":Font",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":TDF",
        ":TDataXtd",
        ":TDocStd",
        ":TNaming",
        ":TPrsStd",
        ":ViewerTest",
    ],
)

cc_library(
    name = "DRAWEXE",
    srcs = glob([
        "DRAWEXE/*.cxx",
    ]),
    hdrs = glob([
        "DRAWEXE/*.gxx",
        "DRAWEXE/*.pxx",
        "DRAWEXE/*.hxx",
        "DRAWEXE/*.lxx",
        "DRAWEXE/*.h",
    ]),
    includes = [
        "DRAWEXE/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Draw",
    ],
)

cc_library(
    name = "Draft",
    srcs = glob([
        "Draft/*.cxx",
    ]),
    hdrs = glob([
        "Draft/*.gxx",
        "Draft/*.pxx",
        "Draft/*.hxx",
        "Draft/*.lxx",
        "Draft/*.h",
    ]),
    includes = [
        "Draft/",
    ],
    deps = [
        ":Adaptor2d",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomProjLib",
        ":IntAna",
        ":IntCurveSurface",
        ":NCollection",
        ":ProjLib",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "Draw",
    srcs = glob([
        "Draw/*.cxx",
    ]),
    hdrs = glob([
        "Draw/*.gxx",
        "Draw/*.pxx",
        "Draw/*.hxx",
        "Draw/*.lxx",
        "Draw/*.h",
    ]),
    includes = [
        "Draw/",
    ],
    deps = [
        ":Aspect",
        ":Bnd",
        ":Cocoa",
        ":ElCLib",
        ":Image",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Plugin",
        ":Resource",
        ":Units",
        ":gp",
        "@tcl",
    ],
)

cc_library(
    name = "DrawDim",
    srcs = glob([
        "DrawDim/*.cxx",
    ]),
    hdrs = glob([
        "DrawDim/*.gxx",
        "DrawDim/*.pxx",
        "DrawDim/*.hxx",
        "DrawDim/*.lxx",
        "DrawDim/*.h",
    ]),
    includes = [
        "DrawDim/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":GC",
        ":Geom2dAPI",
        ":IntAna",
        ":IntAna2d",
        ":NCollection",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "DrawFairCurve",
    srcs = glob([
        "DrawFairCurve/*.cxx",
    ]),
    hdrs = glob([
        "DrawFairCurve/*.gxx",
        "DrawFairCurve/*.pxx",
        "DrawFairCurve/*.hxx",
        "DrawFairCurve/*.lxx",
        "DrawFairCurve/*.h",
    ]),
    includes = [
        "DrawFairCurve/",
    ],
    deps = [
        ":Adaptor2d",
        ":Draw",
        ":FairCurve",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "DsgPrs",
    srcs = glob([
        "DsgPrs/*.cxx",
    ]),
    hdrs = glob([
        "DsgPrs/*.gxx",
        "DsgPrs/*.pxx",
        "DsgPrs/*.hxx",
        "DsgPrs/*.lxx",
        "DsgPrs/*.h",
    ]),
    includes = [
        "DsgPrs/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":ElCLib",
        ":Font",
        ":GC",
        ":IntAna2d",
        ":NCollection",
        ":Quantity",
        ":StdPrs",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Units",
        ":gp",
    ],
)

cc_library(
    name = "ElCLib",
    srcs = glob([
        "ElCLib/*.cxx",
    ]),
    hdrs = glob([
        "ElCLib/*.gxx",
        "ElCLib/*.pxx",
        "ElCLib/*.hxx",
        "ElCLib/*.lxx",
        "ElCLib/*.h",
    ]),
    includes = [
        "ElCLib/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "ElSLib",
    srcs = glob([
        "ElSLib/*.cxx",
    ]),
    hdrs = glob([
        "ElSLib/*.gxx",
        "ElSLib/*.pxx",
        "ElSLib/*.hxx",
        "ElSLib/*.lxx",
        "ElSLib/*.h",
    ]),
    includes = [
        "ElSLib/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "Expr",
    srcs = glob([
        "Expr/*.cxx",
    ]),
    hdrs = glob([
        "Expr/*.gxx",
        "Expr/*.pxx",
        "Expr/*.hxx",
        "Expr/*.lxx",
        "Expr/*.h",
    ]),
    includes = [
        "Expr/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "ExprIntrp",
    srcs = glob([
        "ExprIntrp/*.cxx",
    ]),
    hdrs = glob([
        "ExprIntrp/*.gxx",
        "ExprIntrp/*.pxx",
        "ExprIntrp/*.hxx",
        "ExprIntrp/*.lxx",
        "ExprIntrp/*.h",
    ]),
    includes = [
        "ExprIntrp/",
    ],
    deps = [
        ":Expr",
        ":NCollection",
    ],
)

cc_library(
    name = "FEmTool",
    srcs = glob([
        "FEmTool/*.cxx",
    ]),
    hdrs = glob([
        "FEmTool/*.gxx",
        "FEmTool/*.pxx",
        "FEmTool/*.hxx",
        "FEmTool/*.lxx",
        "FEmTool/*.h",
    ]),
    includes = [
        "FEmTool/",
    ],
    deps = [
        ":GeomAbs",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "FSD",
    srcs = glob([
        "FSD/*.cxx",
    ]),
    hdrs = glob([
        "FSD/*.gxx",
        "FSD/*.pxx",
        "FSD/*.hxx",
        "FSD/*.lxx",
        "FSD/*.h",
    ]),
    includes = [
        "FSD/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
        ":Storage",
    ],
)

cc_library(
    name = "FairCurve",
    srcs = glob([
        "FairCurve/*.cxx",
    ]),
    hdrs = glob([
        "FairCurve/*.gxx",
        "FairCurve/*.pxx",
        "FairCurve/*.hxx",
        "FairCurve/*.lxx",
        "FairCurve/*.h",
    ]),
    includes = [
        "FairCurve/",
    ],
    deps = [
        ":Adaptor2d",
        ":BSplCLib",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "FilletSurf",
    srcs = glob([
        "FilletSurf/*.cxx",
    ]),
    hdrs = glob([
        "FilletSurf/*.gxx",
        "FilletSurf/*.pxx",
        "FilletSurf/*.hxx",
        "FilletSurf/*.lxx",
        "FilletSurf/*.h",
    ]),
    includes = [
        "FilletSurf/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepBlend",
        ":ChFi3d",
        ":ChFiDS",
        ":ElSLib",
        ":GeomAbs",
        ":IntCurveSurface",
        ":NCollection",
        ":StdFail",
        ":TopAbs",
        ":TopOpeBRepDS",
        ":TopTools",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Font",
    srcs = glob([
        "Font/*.cxx",
        "Graphic3d/*.cxx",
        "Prs3d/*.cxx",
        "PrsMgr/*.cxx",
        "Select3D/*.cxx",
        "SelectBasics/*.cxx",
        "SelectMgr/*.cxx",
        "V3d/*.cxx",
    ]),
    hdrs = glob([
        "Font/*.gxx",
        "Font/*.pxx",
        "Font/*.hxx",
        "Font/*.lxx",
        "Font/*.h",
        "Graphic3d/*.gxx",
        "Graphic3d/*.pxx",
        "Graphic3d/*.hxx",
        "Graphic3d/*.lxx",
        "Graphic3d/*.h",
        "Prs3d/*.gxx",
        "Prs3d/*.pxx",
        "Prs3d/*.hxx",
        "Prs3d/*.lxx",
        "Prs3d/*.h",
        "PrsMgr/*.gxx",
        "PrsMgr/*.pxx",
        "PrsMgr/*.hxx",
        "PrsMgr/*.lxx",
        "PrsMgr/*.h",
        "Select3D/*.gxx",
        "Select3D/*.pxx",
        "Select3D/*.hxx",
        "Select3D/*.lxx",
        "Select3D/*.h",
        "SelectBasics/*.gxx",
        "SelectBasics/*.pxx",
        "SelectBasics/*.hxx",
        "SelectBasics/*.lxx",
        "SelectBasics/*.h",
        "SelectMgr/*.gxx",
        "SelectMgr/*.pxx",
        "SelectMgr/*.hxx",
        "SelectMgr/*.lxx",
        "SelectMgr/*.h",
        "V3d/*.gxx",
        "V3d/*.pxx",
        "V3d/*.hxx",
        "V3d/*.lxx",
        "V3d/*.h",
    ]),
    includes = [
        "Font/",
        "Graphic3d/",
        "Prs3d/",
        "PrsMgr/",
        "Select3D/",
        "SelectBasics/",
        "SelectMgr/",
        "V3d/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":BVH",
        ":Bnd",
        ":GC",
        ":GCE2d",
        ":GeomLib",
        ":HLRAlgo",
        ":Image",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
        "@freetype2",
    ],
)

cc_library(
    name = "GC",
    srcs = glob([
        "GC/*.cxx",
    ]),
    hdrs = glob([
        "GC/*.gxx",
        "GC/*.pxx",
        "GC/*.hxx",
        "GC/*.lxx",
        "GC/*.h",
    ]),
    includes = [
        "GC/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "GCE2d",
    srcs = glob([
        "GCE2d/*.cxx",
    ]),
    hdrs = glob([
        "GCE2d/*.gxx",
        "GCE2d/*.pxx",
        "GCE2d/*.hxx",
        "GCE2d/*.lxx",
        "GCE2d/*.h",
    ]),
    includes = [
        "GCE2d/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":IntAna2d",
        ":NCollection",
        ":StdFail",
        ":gp",
    ],
)

cc_library(
    name = "GProp",
    srcs = glob([
        "GProp/*.cxx",
    ]),
    hdrs = glob([
        "GProp/*.gxx",
        "GProp/*.pxx",
        "GProp/*.hxx",
        "GProp/*.lxx",
        "GProp/*.h",
    ]),
    includes = [
        "GProp/",
    ],
    deps = [
        ":ElCLib",
        ":NCollection",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GccAna",
    srcs = glob([
        "GccAna/*.cxx",
    ]),
    hdrs = glob([
        "GccAna/*.gxx",
        "GccAna/*.pxx",
        "GccAna/*.hxx",
        "GccAna/*.lxx",
        "GccAna/*.h",
    ]),
    includes = [
        "GccAna/",
    ],
    deps = [
        ":ElCLib",
        ":GccEnt",
        ":GccInt",
        ":IntAna2d",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GccEnt",
    srcs = glob([
        "GccEnt/*.cxx",
    ]),
    hdrs = glob([
        "GccEnt/*.gxx",
        "GccEnt/*.pxx",
        "GccEnt/*.hxx",
        "GccEnt/*.lxx",
        "GccEnt/*.h",
    ]),
    includes = [
        "GccEnt/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "GccInt",
    srcs = glob([
        "GccInt/*.cxx",
    ]),
    hdrs = glob([
        "GccInt/*.gxx",
        "GccInt/*.pxx",
        "GccInt/*.hxx",
        "GccInt/*.lxx",
        "GccInt/*.h",
    ]),
    includes = [
        "GccInt/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "Geom2dAPI",
    srcs = glob([
        "Geom2dAPI/*.cxx",
    ]),
    hdrs = glob([
        "Geom2dAPI/*.gxx",
        "Geom2dAPI/*.pxx",
        "Geom2dAPI/*.hxx",
        "Geom2dAPI/*.lxx",
        "Geom2dAPI/*.h",
    ]),
    includes = [
        "Geom2dAPI/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":BSplCLib",
        ":Geom2dInt",
        ":GeomAbs",
        ":IntRes2d",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Geom2dGcc",
    srcs = glob([
        "Geom2dGcc/*.cxx",
    ]),
    hdrs = glob([
        "Geom2dGcc/*.gxx",
        "Geom2dGcc/*.pxx",
        "Geom2dGcc/*.hxx",
        "Geom2dGcc/*.lxx",
        "Geom2dGcc/*.h",
    ]),
    includes = [
        "Geom2dGcc/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":GccAna",
        ":GccEnt",
        ":GccInt",
        ":Geom2dInt",
        ":IntAna2d",
        ":IntRes2d",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Geom2dHatch",
    srcs = glob([
        "Geom2dHatch/*.cxx",
    ]),
    hdrs = glob([
        "Geom2dHatch/*.gxx",
        "Geom2dHatch/*.pxx",
        "Geom2dHatch/*.hxx",
        "Geom2dHatch/*.lxx",
        "Geom2dHatch/*.h",
    ]),
    includes = [
        "Geom2dHatch/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":Geom2dInt",
        ":HatchGen",
        ":IntRes2d",
        ":NCollection",
        ":StdFail",
        ":TopAbs",
        ":TopClass",
        ":TopTrans",
        ":gp",
    ],
)

cc_library(
    name = "Geom2dInt",
    srcs = glob([
        "Geom2dInt/*.cxx",
    ]),
    hdrs = glob([
        "Geom2dInt/*.gxx",
        "Geom2dInt/*.pxx",
        "Geom2dInt/*.hxx",
        "Geom2dInt/*.lxx",
        "Geom2dInt/*.h",
    ]),
    includes = [
        "Geom2dInt/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":GeomAbs",
        ":IntCurve",
        ":IntImpParGen",
        ":IntRes2d",
        ":Intf",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Geom2dToIGES",
    srcs = glob([
        "Geom2dToIGES/*.cxx",
    ]),
    hdrs = glob([
        "Geom2dToIGES/*.gxx",
        "Geom2dToIGES/*.pxx",
        "Geom2dToIGES/*.hxx",
        "Geom2dToIGES/*.lxx",
        "Geom2dToIGES/*.h",
    ]),
    includes = [
        "Geom2dToIGES/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":GeomToIGES",
        ":IGESBasic",
        ":Interface",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "GeomAbs",
    srcs = glob([
        "GeomAbs/*.cxx",
    ]),
    hdrs = glob([
        "GeomAbs/*.gxx",
        "GeomAbs/*.pxx",
        "GeomAbs/*.hxx",
        "GeomAbs/*.lxx",
        "GeomAbs/*.h",
    ]),
    includes = [
        "GeomAbs/",
    ],
    deps = [
    ],
)

cc_library(
    name = "GeomFill",
    srcs = glob([
        "GeomFill/*.cxx",
    ]),
    hdrs = glob([
        "GeomFill/*.gxx",
        "GeomFill/*.pxx",
        "GeomFill/*.hxx",
        "GeomFill/*.lxx",
        "GeomFill/*.h",
    ]),
    includes = [
        "GeomFill/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":AppBlend",
        ":Approx",
        ":BSplCLib",
        ":Bnd",
        ":CSLib",
        ":Convert",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":GeomLib",
        ":IntCurveSurface",
        ":Law",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GeomLib",
    srcs = glob([
        "GeomLib/*.cxx",
    ]),
    hdrs = glob([
        "GeomLib/*.gxx",
        "GeomLib/*.pxx",
        "GeomLib/*.hxx",
        "GeomLib/*.lxx",
        "GeomLib/*.h",
    ]),
    includes = [
        "GeomLib/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":BSplCLib",
        ":BSplSLib",
        ":CSLib",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GeomPlate",
    srcs = glob([
        "GeomPlate/*.cxx",
    ]),
    hdrs = glob([
        "GeomPlate/*.gxx",
        "GeomPlate/*.pxx",
        "GeomPlate/*.hxx",
        "GeomPlate/*.lxx",
        "GeomPlate/*.h",
    ]),
    includes = [
        "GeomPlate/",
    ],
    deps = [
        ":Adaptor2d",
        ":AdvApprox",
        ":Approx",
        ":ApproxInt",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomLib",
        ":IntRes2d",
        ":Law",
        ":LocalAnalysis",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":Plate",
        ":ProjLib",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GeomProjLib",
    srcs = glob([
        "GeomProjLib/*.cxx",
    ]),
    hdrs = glob([
        "GeomProjLib/*.gxx",
        "GeomProjLib/*.pxx",
        "GeomProjLib/*.hxx",
        "GeomProjLib/*.lxx",
        "GeomProjLib/*.h",
    ]),
    includes = [
        "GeomProjLib/",
    ],
    deps = [
        ":Adaptor2d",
        ":Approx",
        ":NCollection",
        ":ProjLib",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "GeomToIGES",
    srcs = glob([
        "GeomToIGES/*.cxx",
    ]),
    hdrs = glob([
        "GeomToIGES/*.gxx",
        "GeomToIGES/*.pxx",
        "GeomToIGES/*.hxx",
        "GeomToIGES/*.lxx",
        "GeomToIGES/*.h",
    ]),
    includes = [
        "GeomToIGES/",
    ],
    deps = [
        ":Adaptor2d",
        ":BSplCLib",
        ":IGESBasic",
        ":IGESConvGeom",
        ":IGESSolid",
        ":Interface",
        ":NCollection",
        ":ShapeAnalysis",
        ":ShapeCustom",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "GeomToStep",
    srcs = glob([
        "GeomToStep/*.cxx",
    ]),
    hdrs = glob([
        "GeomToStep/*.gxx",
        "GeomToStep/*.pxx",
        "GeomToStep/*.hxx",
        "GeomToStep/*.lxx",
        "GeomToStep/*.h",
    ]),
    includes = [
        "GeomToStep/",
    ],
    deps = [
        ":Adaptor2d",
        ":GeomAbs",
        ":NCollection",
        ":StdFail",
        ":StepData",
        ":StepGeom",
        ":TColgp",
        ":UnitsMethods",
        ":gp",
    ],
)

cc_library(
    name = "GeometryTest",
    srcs = glob([
        "GeometryTest/*.cxx",
    ]),
    hdrs = glob([
        "GeometryTest/*.gxx",
        "GeometryTest/*.pxx",
        "GeometryTest/*.hxx",
        "GeometryTest/*.lxx",
        "GeometryTest/*.h",
    ]),
    includes = [
        "GeometryTest/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":ApproxInt",
        ":BSplCLib",
        ":Draw",
        ":DrawFairCurve",
        ":FairCurve",
        ":GC",
        ":GccAna",
        ":GccEnt",
        ":GccInt",
        ":Geom2dAPI",
        ":Geom2dGcc",
        ":GeomAbs",
        ":GeomFill",
        ":GeomProjLib",
        ":GeomliteTest",
        ":Law",
        ":LocalAnalysis",
        ":NCollection",
        ":Poly",
        ":ProjLib",
        ":TColgp",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "GeomliteTest",
    srcs = glob([
        "GeomliteTest/*.cxx",
    ]),
    hdrs = glob([
        "GeomliteTest/*.gxx",
        "GeomliteTest/*.pxx",
        "GeomliteTest/*.hxx",
        "GeomliteTest/*.lxx",
        "GeomliteTest/*.h",
    ]),
    includes = [
        "GeomliteTest/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppDef",
        ":AppParCurves",
        ":Approx",
        ":ApproxInt",
        ":BSplCLib",
        ":Convert",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomLib",
        ":IntAna2d",
        ":IntRes2d",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "HLRAlgo",
    srcs = glob([
        "HLRAlgo/*.cxx",
    ]),
    hdrs = glob([
        "HLRAlgo/*.gxx",
        "HLRAlgo/*.pxx",
        "HLRAlgo/*.hxx",
        "HLRAlgo/*.lxx",
        "HLRAlgo/*.h",
    ]),
    includes = [
        "HLRAlgo/",
    ],
    deps = [
        ":Intrv",
        ":NCollection",
        ":TColgp",
        ":TopAbs",
        ":TopBas",
        ":gp",
    ],
)

cc_library(
    name = "HLRAppli",
    srcs = glob([
        "HLRAppli/*.cxx",
    ]),
    hdrs = glob([
        "HLRAppli/*.gxx",
        "HLRAppli/*.pxx",
        "HLRAppli/*.hxx",
        "HLRAppli/*.lxx",
        "HLRAppli/*.h",
    ]),
    includes = [
        "HLRAppli/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":HLRAlgo",
        ":NCollection",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "HLRTest",
    srcs = glob([
        "HLRTest/*.cxx",
    ]),
    hdrs = glob([
        "HLRTest/*.gxx",
        "HLRTest/*.pxx",
        "HLRTest/*.hxx",
        "HLRTest/*.lxx",
        "HLRTest/*.h",
    ]),
    includes = [
        "HLRTest/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":Draw",
        ":HLRAlgo",
        ":HLRAppli",
        ":NCollection",
        ":OSD",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "Hatch",
    srcs = glob([
        "Hatch/*.cxx",
    ]),
    hdrs = glob([
        "Hatch/*.gxx",
        "Hatch/*.pxx",
        "Hatch/*.hxx",
        "Hatch/*.lxx",
        "Hatch/*.h",
    ]),
    includes = [
        "Hatch/",
    ],
    deps = [
        ":IntAna2d",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "HatchGen",
    srcs = glob([
        "HatchGen/*.cxx",
    ]),
    hdrs = glob([
        "HatchGen/*.gxx",
        "HatchGen/*.pxx",
        "HatchGen/*.hxx",
        "HatchGen/*.lxx",
        "HatchGen/*.h",
    ]),
    includes = [
        "HatchGen/",
    ],
    deps = [
        ":IntRes2d",
        ":NCollection",
        ":TopAbs",
    ],
)

cc_library(
    name = "HeaderSection",
    srcs = glob([
        "HeaderSection/*.cxx",
    ]),
    hdrs = glob([
        "HeaderSection/*.gxx",
        "HeaderSection/*.pxx",
        "HeaderSection/*.hxx",
        "HeaderSection/*.lxx",
        "HeaderSection/*.h",
    ]),
    includes = [
        "HeaderSection/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepData",
    ],
)

cc_library(
    name = "IFGraph",
    srcs = glob([
        "IFGraph/*.cxx",
    ]),
    hdrs = glob([
        "IFGraph/*.gxx",
        "IFGraph/*.pxx",
        "IFGraph/*.hxx",
        "IFGraph/*.lxx",
        "IFGraph/*.h",
    ]),
    includes = [
        "IFGraph/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
    ],
)

cc_library(
    name = "IFSelect",
    srcs = glob([
        "IFSelect/*.cxx",
    ]),
    hdrs = glob([
        "IFSelect/*.gxx",
        "IFSelect/*.pxx",
        "IFSelect/*.hxx",
        "IFSelect/*.lxx",
        "IFSelect/*.h",
    ]),
    includes = [
        "IFSelect/",
    ],
    deps = [
        ":IFGraph",
        ":Interface",
        ":Message",
        ":MoniTool",
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "IGESAppli",
    srcs = glob([
        "IGESAppli/*.cxx",
    ]),
    hdrs = glob([
        "IGESAppli/*.gxx",
        "IGESAppli/*.pxx",
        "IGESAppli/*.hxx",
        "IGESAppli/*.lxx",
        "IGESAppli/*.h",
    ]),
    includes = [
        "IGESAppli/",
    ],
    deps = [
        ":IGESBasic",
        ":IGESDefs",
        ":IGESDimen",
        ":IGESDraw",
        ":IGESGraph",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESBasic",
    srcs = glob([
        "IGESBasic/*.cxx",
        "IGESData/*.cxx",
        "IGESGeom/*.cxx",
    ]),
    hdrs = glob([
        "IGESBasic/*.gxx",
        "IGESBasic/*.pxx",
        "IGESBasic/*.hxx",
        "IGESBasic/*.lxx",
        "IGESBasic/*.h",
        "IGESData/*.gxx",
        "IGESData/*.pxx",
        "IGESData/*.hxx",
        "IGESData/*.lxx",
        "IGESData/*.h",
        "IGESGeom/*.gxx",
        "IGESGeom/*.pxx",
        "IGESGeom/*.hxx",
        "IGESGeom/*.lxx",
        "IGESGeom/*.h",
    ]),
    includes = [
        "IGESBasic/",
        "IGESData/",
        "IGESGeom/",
    ],
    deps = [
        ":Interface",
        ":LibCtl",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":TColgp",
        ":UnitsMethods",
        ":gp",
    ],
)

cc_library(
    name = "IGESCAFControl",
    srcs = glob([
        "IGESCAFControl/*.cxx",
    ]),
    hdrs = glob([
        "IGESCAFControl/*.gxx",
        "IGESCAFControl/*.pxx",
        "IGESCAFControl/*.hxx",
        "IGESCAFControl/*.lxx",
        "IGESCAFControl/*.h",
    ]),
    includes = [
        "IGESCAFControl/",
    ],
    deps = [
        ":BRep",
        ":IGESBasic",
        ":IGESControl",
        ":IGESGraph",
        ":IGESSolid",
        ":Interface",
        ":NCollection",
        ":Quantity",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XCAFDimTolObjects",
        ":XCAFPrs",
        ":XSControl",
    ],
)

cc_library(
    name = "IGESControl",
    srcs = glob([
        "IGESControl/*.cxx",
    ]),
    hdrs = glob([
        "IGESControl/*.gxx",
        "IGESControl/*.pxx",
        "IGESControl/*.hxx",
        "IGESControl/*.lxx",
        "IGESControl/*.h",
    ]),
    includes = [
        "IGESControl/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepToIGES",
        ":BRepToIGESBRep",
        ":Bnd",
        ":GeomToIGES",
        ":IFSelect",
        ":IGESAppli",
        ":IGESBasic",
        ":IGESSelect",
        ":IGESSolid",
        ":IGESToBRep",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeExtend",
        ":TopExp",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XSAlgo",
        ":XSControl",
        ":gp",
    ],
)

cc_library(
    name = "IGESConvGeom",
    srcs = glob([
        "IGESConvGeom/*.cxx",
    ]),
    hdrs = glob([
        "IGESConvGeom/*.gxx",
        "IGESConvGeom/*.pxx",
        "IGESConvGeom/*.hxx",
        "IGESConvGeom/*.lxx",
        "IGESConvGeom/*.h",
    ]),
    includes = [
        "IGESConvGeom/",
    ],
    deps = [
        ":Adaptor2d",
        ":BSplCLib",
        ":BSplSLib",
        ":IGESBasic",
        ":Interface",
        ":NCollection",
        ":PLib",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESDefs",
    srcs = glob([
        "IGESDefs/*.cxx",
    ]),
    hdrs = glob([
        "IGESDefs/*.gxx",
        "IGESDefs/*.pxx",
        "IGESDefs/*.hxx",
        "IGESDefs/*.lxx",
        "IGESDefs/*.h",
    ]),
    includes = [
        "IGESDefs/",
    ],
    deps = [
        ":IGESBasic",
        ":IGESGraph",
        ":Interface",
        ":Message",
        ":NCollection",
    ],
)

cc_library(
    name = "IGESDimen",
    srcs = glob([
        "IGESDimen/*.cxx",
    ]),
    hdrs = glob([
        "IGESDimen/*.gxx",
        "IGESDimen/*.pxx",
        "IGESDimen/*.hxx",
        "IGESDimen/*.lxx",
        "IGESDimen/*.h",
    ]),
    includes = [
        "IGESDimen/",
    ],
    deps = [
        ":IGESBasic",
        ":IGESGraph",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESDraw",
    srcs = glob([
        "IGESDraw/*.cxx",
    ]),
    hdrs = glob([
        "IGESDraw/*.gxx",
        "IGESDraw/*.pxx",
        "IGESDraw/*.hxx",
        "IGESDraw/*.lxx",
        "IGESDraw/*.h",
    ]),
    includes = [
        "IGESDraw/",
    ],
    deps = [
        ":IGESBasic",
        ":IGESDimen",
        ":IGESGraph",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESFile",
    srcs = glob([
        "IGESFile/*.cxx",
    ]),
    hdrs = glob([
        "IGESFile/*.gxx",
        "IGESFile/*.pxx",
        "IGESFile/*.hxx",
        "IGESFile/*.lxx",
        "IGESFile/*.h",
    ]),
    includes = [
        "IGESFile/",
    ],
    deps = [
        ":IGESBasic",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "IGESGraph",
    srcs = glob([
        "IGESGraph/*.cxx",
    ]),
    hdrs = glob([
        "IGESGraph/*.gxx",
        "IGESGraph/*.pxx",
        "IGESGraph/*.hxx",
        "IGESGraph/*.lxx",
        "IGESGraph/*.h",
    ]),
    includes = [
        "IGESGraph/",
    ],
    deps = [
        ":IGESBasic",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESSelect",
    srcs = glob([
        "IGESSelect/*.cxx",
    ]),
    hdrs = glob([
        "IGESSelect/*.gxx",
        "IGESSelect/*.pxx",
        "IGESSelect/*.hxx",
        "IGESSelect/*.lxx",
        "IGESSelect/*.h",
    ]),
    includes = [
        "IGESSelect/",
    ],
    deps = [
        ":IFGraph",
        ":IFSelect",
        ":IGESAppli",
        ":IGESBasic",
        ":IGESDefs",
        ":IGESDraw",
        ":IGESFile",
        ":IGESGraph",
        ":IGESSolid",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESSolid",
    srcs = glob([
        "IGESSolid/*.cxx",
    ]),
    hdrs = glob([
        "IGESSolid/*.gxx",
        "IGESSolid/*.pxx",
        "IGESSolid/*.hxx",
        "IGESSolid/*.lxx",
        "IGESSolid/*.h",
    ]),
    includes = [
        "IGESSolid/",
    ],
    deps = [
        ":IGESBasic",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IGESToBRep",
    srcs = glob([
        "IGESToBRep/*.cxx",
    ]),
    hdrs = glob([
        "IGESToBRep/*.gxx",
        "IGESToBRep/*.pxx",
        "IGESToBRep/*.hxx",
        "IGESToBRep/*.lxx",
        "IGESToBRep/*.h",
    ]),
    includes = [
        "IGESToBRep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepFill",
        ":BRepPrimAPI",
        ":BSplCLib",
        ":ElCLib",
        ":ElSLib",
        ":GProp",
        ":GeomAbs",
        ":GeomLib",
        ":IGESAppli",
        ":IGESBasic",
        ":IGESConvGeom",
        ":IGESFile",
        ":IGESSolid",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XSAlgo",
        ":gp",
    ],
)

cc_library(
    name = "IVtk",
    srcs = glob([
        "IVtk/*.cxx",
    ]),
    hdrs = glob([
        "IVtk/*.gxx",
        "IVtk/*.pxx",
        "IVtk/*.hxx",
        "IVtk/*.lxx",
        "IVtk/*.h",
    ]),
    includes = [
        "IVtk/",
    ],
    deps = [
        ":Font",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "IVtkDraw",
    srcs = glob([
        "IVtkDraw/*.cxx",
    ]),
    hdrs = glob([
        "IVtkDraw/*.gxx",
        "IVtkDraw/*.pxx",
        "IVtkDraw/*.hxx",
        "IVtkDraw/*.lxx",
        "IVtkDraw/*.h",
    ]),
    includes = [
        "IVtkDraw/",
    ],
    deps = [
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Draw",
        ":IVtk",
        ":IVtkOCC",
        ":IVtkTools",
        ":IVtkVTK",
        ":Message",
        ":NCollection",
        ":OpenGl",
        ":TopTools",
        ":TopoDS",
        ":WNT",
        ":Xw",
    ],
)

cc_library(
    name = "IVtkOCC",
    srcs = glob([
        "IVtkOCC/*.cxx",
    ]),
    hdrs = glob([
        "IVtkOCC/*.gxx",
        "IVtkOCC/*.pxx",
        "IVtkOCC/*.hxx",
        "IVtkOCC/*.lxx",
        "IVtkOCC/*.h",
    ]),
    includes = [
        "IVtkOCC/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Bnd",
        ":Font",
        ":Hatch",
        ":IVtk",
        ":Message",
        ":NCollection",
        ":Poly",
        ":StdSelect",
        ":TColgp",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "IVtkTools",
    srcs = glob([
        "IVtkTools/*.cxx",
    ]),
    hdrs = glob([
        "IVtkTools/*.gxx",
        "IVtkTools/*.pxx",
        "IVtkTools/*.hxx",
        "IVtkTools/*.lxx",
        "IVtkTools/*.h",
    ]),
    includes = [
        "IVtkTools/",
    ],
    deps = [
        ":IVtk",
        ":IVtkOCC",
        ":IVtkVTK",
        ":NCollection",
    ],
)

cc_library(
    name = "IVtkVTK",
    srcs = glob([
        "IVtkVTK/*.cxx",
    ]),
    hdrs = glob([
        "IVtkVTK/*.gxx",
        "IVtkVTK/*.pxx",
        "IVtkVTK/*.hxx",
        "IVtkVTK/*.lxx",
        "IVtkVTK/*.h",
    ]),
    includes = [
        "IVtkVTK/",
    ],
    deps = [
        ":IVtk",
    ],
)

cc_library(
    name = "Image",
    srcs = glob([
        "Image/*.cxx",
    ]),
    hdrs = glob([
        "Image/*.gxx",
        "Image/*.pxx",
        "Image/*.hxx",
        "Image/*.lxx",
        "Image/*.h",
    ]),
    includes = [
        "Image/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":Resource",
        ":gp",
    ],
)

cc_library(
    name = "IntAna",
    srcs = glob([
        "IntAna/*.cxx",
    ]),
    hdrs = glob([
        "IntAna/*.gxx",
        "IntAna/*.pxx",
        "IntAna/*.hxx",
        "IntAna/*.lxx",
        "IntAna/*.h",
    ]),
    includes = [
        "IntAna/",
    ],
    deps = [
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":IntAna2d",
        ":NCollection",
        ":StdFail",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntAna2d",
    srcs = glob([
        "IntAna2d/*.cxx",
    ]),
    hdrs = glob([
        "IntAna2d/*.gxx",
        "IntAna2d/*.pxx",
        "IntAna2d/*.hxx",
        "IntAna2d/*.lxx",
        "IntAna2d/*.h",
    ]),
    includes = [
        "IntAna2d/",
    ],
    deps = [
        ":ElCLib",
        ":NCollection",
        ":StdFail",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntCurve",
    srcs = glob([
        "IntCurve/*.cxx",
    ]),
    hdrs = glob([
        "IntCurve/*.gxx",
        "IntCurve/*.pxx",
        "IntCurve/*.hxx",
        "IntCurve/*.lxx",
        "IntCurve/*.h",
    ]),
    includes = [
        "IntCurve/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":ElCLib",
        ":GeomAbs",
        ":IntAna2d",
        ":IntImpParGen",
        ":IntRes2d",
        ":Intf",
        ":NCollection",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntCurveSurface",
    srcs = glob([
        "IntCurveSurface/*.cxx",
    ]),
    hdrs = glob([
        "IntCurveSurface/*.gxx",
        "IntCurveSurface/*.pxx",
        "IntCurveSurface/*.hxx",
        "IntCurveSurface/*.lxx",
        "IntCurveSurface/*.h",
    ]),
    includes = [
        "IntCurveSurface/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":ElCLib",
        ":ElSLib",
        ":GProp",
        ":GeomAbs",
        ":IntAna",
        ":IntAna2d",
        ":IntImp",
        ":IntSurf",
        ":Intf",
        ":NCollection",
        ":ProjLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntImp",
    srcs = glob([
        "IntImp/*.cxx",
    ]),
    hdrs = glob([
        "IntImp/*.gxx",
        "IntImp/*.pxx",
        "IntImp/*.hxx",
        "IntImp/*.lxx",
        "IntImp/*.h",
    ]),
    includes = [
        "IntImp/",
    ],
    deps = [
        ":NCollection",
        ":StdFail",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntImpParGen",
    srcs = glob([
        "IntImpParGen/*.cxx",
    ]),
    hdrs = glob([
        "IntImpParGen/*.gxx",
        "IntImpParGen/*.pxx",
        "IntImpParGen/*.hxx",
        "IntImpParGen/*.lxx",
        "IntImpParGen/*.h",
    ]),
    includes = [
        "IntImpParGen/",
    ],
    deps = [
        ":IntRes2d",
        ":NCollection",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "IntPolyh",
    srcs = glob([
        "IntPolyh/*.cxx",
    ]),
    hdrs = glob([
        "IntPolyh/*.gxx",
        "IntPolyh/*.pxx",
        "IntPolyh/*.hxx",
        "IntPolyh/*.lxx",
        "IntPolyh/*.h",
    ]),
    includes = [
        "IntPolyh/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":IntCurveSurface",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "IntRes2d",
    srcs = glob([
        "IntRes2d/*.cxx",
    ]),
    hdrs = glob([
        "IntRes2d/*.gxx",
        "IntRes2d/*.pxx",
        "IntRes2d/*.hxx",
        "IntRes2d/*.lxx",
        "IntRes2d/*.h",
    ]),
    includes = [
        "IntRes2d/",
    ],
    deps = [
        ":NCollection",
        ":StdFail",
        ":gp",
    ],
)

cc_library(
    name = "IntSurf",
    srcs = glob([
        "IntSurf/*.cxx",
    ]),
    hdrs = glob([
        "IntSurf/*.gxx",
        "IntSurf/*.pxx",
        "IntSurf/*.hxx",
        "IntSurf/*.lxx",
        "IntSurf/*.h",
    ]),
    includes = [
        "IntSurf/",
    ],
    deps = [
        ":Adaptor2d",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":NCollection",
        ":StdFail",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "IntTools",
    srcs = glob([
        "IntTools/*.cxx",
    ]),
    hdrs = glob([
        "IntTools/*.gxx",
        "IntTools/*.pxx",
        "IntTools/*.hxx",
        "IntTools/*.lxx",
        "IntTools/*.h",
    ]),
    includes = [
        "IntTools/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Bnd",
        ":CSLib",
        ":ElCLib",
        ":ElSLib",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dHatch",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomLib",
        ":IntAna",
        ":IntCurveSurface",
        ":IntRes2d",
        ":IntSurf",
        ":NCollection",
        ":ProjLib",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "IntWalk",
    srcs = glob([
        "IntWalk/*.cxx",
    ]),
    hdrs = glob([
        "IntWalk/*.gxx",
        "IntWalk/*.pxx",
        "IntWalk/*.hxx",
        "IntWalk/*.lxx",
        "IntWalk/*.h",
    ]),
    includes = [
        "IntWalk/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":IntImp",
        ":IntSurf",
        ":NCollection",
        ":OSD",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Interface",
    srcs = glob([
        "Interface/*.cxx",
    ]),
    hdrs = glob([
        "Interface/*.gxx",
        "Interface/*.pxx",
        "Interface/*.hxx",
        "Interface/*.lxx",
        "Interface/*.h",
    ]),
    includes = [
        "Interface/",
    ],
    deps = [
        ":LibCtl",
        ":Message",
        ":MoniTool",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":XSMessage",
    ],
)

cc_library(
    name = "InterfaceGraphic",
    srcs = glob([
        "InterfaceGraphic/*.cxx",
    ]),
    hdrs = glob([
        "InterfaceGraphic/*.gxx",
        "InterfaceGraphic/*.pxx",
        "InterfaceGraphic/*.hxx",
        "InterfaceGraphic/*.lxx",
        "InterfaceGraphic/*.h",
    ]),
    includes = [
        "InterfaceGraphic/",
    ],
    deps = [
    ],
)

cc_library(
    name = "Intf",
    srcs = glob([
        "Intf/*.cxx",
    ]),
    hdrs = glob([
        "Intf/*.gxx",
        "Intf/*.pxx",
        "Intf/*.hxx",
        "Intf/*.lxx",
        "Intf/*.h",
    ]),
    includes = [
        "Intf/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":ElCLib",
        ":IntAna",
        ":IntAna2d",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "Intrv",
    srcs = glob([
        "Intrv/*.cxx",
    ]),
    hdrs = glob([
        "Intrv/*.gxx",
        "Intrv/*.pxx",
        "Intrv/*.hxx",
        "Intrv/*.lxx",
        "Intrv/*.h",
    ]),
    includes = [
        "Intrv/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "LDOM",
    srcs = glob([
        "LDOM/*.cxx",
    ]),
    hdrs = glob([
        "LDOM/*.gxx",
        "LDOM/*.pxx",
        "LDOM/*.hxx",
        "LDOM/*.lxx",
        "LDOM/*.h",
    ]),
    includes = [
        "LDOM/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "LProp",
    srcs = glob([
        "LProp/*.cxx",
    ]),
    hdrs = glob([
        "LProp/*.gxx",
        "LProp/*.pxx",
        "LProp/*.hxx",
        "LProp/*.lxx",
        "LProp/*.h",
    ]),
    includes = [
        "LProp/",
    ],
    deps = [
        ":CSLib",
        ":ElCLib",
        ":GeomAbs",
        ":NCollection",
        ":TColgp",
        ":math",
    ],
)

cc_library(
    name = "LProp3d",
    srcs = glob([
        "LProp3d/*.cxx",
    ]),
    hdrs = glob([
        "LProp3d/*.gxx",
        "LProp3d/*.pxx",
        "LProp3d/*.hxx",
        "LProp3d/*.lxx",
        "LProp3d/*.h",
    ]),
    includes = [
        "LProp3d/",
    ],
    deps = [
        ":Adaptor2d",
        ":LProp",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "Law",
    srcs = glob([
        "Law/*.cxx",
    ]),
    hdrs = glob([
        "Law/*.gxx",
        "Law/*.pxx",
        "Law/*.hxx",
        "Law/*.lxx",
        "Law/*.h",
    ]),
    includes = [
        "Law/",
    ],
    deps = [
        ":Adaptor2d",
        ":BSplCLib",
        ":Draw",
        ":ElCLib",
        ":GeomAbs",
        ":NCollection",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "LibCtl",
    srcs = glob([
        "LibCtl/*.cxx",
    ]),
    hdrs = glob([
        "LibCtl/*.gxx",
        "LibCtl/*.pxx",
        "LibCtl/*.hxx",
        "LibCtl/*.lxx",
        "LibCtl/*.h",
    ]),
    includes = [
        "LibCtl/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "LocOpe",
    srcs = glob([
        "LocOpe/*.cxx",
    ]),
    hdrs = glob([
        "LocOpe/*.gxx",
        "LocOpe/*.pxx",
        "LocOpe/*.hxx",
        "LocOpe/*.lxx",
        "LocOpe/*.h",
    ]),
    includes = [
        "LocOpe/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepAlgo",
        ":BRepFill",
        ":BRepIntCurveSurface",
        ":BRepSweep",
        ":BSplCLib",
        ":Bnd",
        ":ElCLib",
        ":ElSLib",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dInt",
        ":GeomFill",
        ":GeomProjLib",
        ":IntAna",
        ":IntCurveSurface",
        ":NCollection",
        ":ShapeAnalysis",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "LocalAnalysis",
    srcs = glob([
        "LocalAnalysis/*.cxx",
    ]),
    hdrs = glob([
        "LocalAnalysis/*.gxx",
        "LocalAnalysis/*.pxx",
        "LocalAnalysis/*.hxx",
        "LocalAnalysis/*.lxx",
        "LocalAnalysis/*.h",
    ]),
    includes = [
        "LocalAnalysis/",
    ],
    deps = [
        ":Adaptor2d",
        ":GeomAbs",
        ":NCollection",
        ":StdFail",
        ":gp",
    ],
)

cc_library(
    name = "MAT",
    srcs = glob([
        "MAT/*.cxx",
    ]),
    hdrs = glob([
        "MAT/*.gxx",
        "MAT/*.pxx",
        "MAT/*.hxx",
        "MAT/*.lxx",
        "MAT/*.h",
    ]),
    includes = [
        "MAT/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "MAT2d",
    srcs = glob([
        "MAT2d/*.cxx",
    ]),
    hdrs = glob([
        "MAT2d/*.gxx",
        "MAT2d/*.pxx",
        "MAT2d/*.hxx",
        "MAT2d/*.lxx",
        "MAT2d/*.h",
    ]),
    includes = [
        "MAT2d/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Bisector",
        ":Draw",
        ":ElCLib",
        ":GCE2d",
        ":Geom2dInt",
        ":GeomAbs",
        ":IntRes2d",
        ":MAT",
        ":NCollection",
        ":StdFail",
        ":gp",
    ],
)

cc_library(
    name = "MMgt",
    srcs = glob([
        "MMgt/*.cxx",
    ]),
    hdrs = glob([
        "MMgt/*.gxx",
        "MMgt/*.pxx",
        "MMgt/*.hxx",
        "MMgt/*.lxx",
        "MMgt/*.h",
    ]),
    includes = [
        "MMgt/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "MeshTest",
    srcs = glob([
        "MeshTest/*.cxx",
    ]),
    hdrs = glob([
        "MeshTest/*.gxx",
        "MeshTest/*.pxx",
        "MeshTest/*.hxx",
        "MeshTest/*.lxx",
        "MeshTest/*.h",
    ]),
    includes = [
        "MeshTest/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppCont",
        ":ApproxInt",
        ":BRep",
        ":BRepTest",
        ":Bnd",
        ":CSLib",
        ":Draw",
        ":GProp",
        ":GeometryTest",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":Poly",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "MeshVS",
    srcs = glob([
        "MeshVS/*.cxx",
    ]),
    hdrs = glob([
        "MeshVS/*.gxx",
        "MeshVS/*.pxx",
        "MeshVS/*.hxx",
        "MeshVS/*.lxx",
        "MeshVS/*.h",
    ]),
    includes = [
        "MeshVS/",
    ],
    deps = [
        ":AIS",
        ":Aspect",
        ":Bnd",
        ":Font",
        ":Image",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":StdSelect",
        ":TColgp",
        ":TopLoc",
        ":gp",
    ],
)

cc_library(
    name = "Message",
    srcs = glob([
        "Message/*.cxx",
    ]),
    hdrs = glob([
        "Message/*.gxx",
        "Message/*.pxx",
        "Message/*.hxx",
        "Message/*.lxx",
        "Message/*.h",
    ]),
    includes = [
        "Message/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "MoniTool",
    srcs = glob([
        "MoniTool/*.cxx",
    ]),
    hdrs = glob([
        "MoniTool/*.gxx",
        "MoniTool/*.pxx",
        "MoniTool/*.hxx",
        "MoniTool/*.lxx",
        "MoniTool/*.h",
    ]),
    includes = [
        "MoniTool/",
    ],
    deps = [
        ":Adaptor2d",
        ":Message",
        ":NCollection",
        ":OSD",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "NCollection",
    srcs = glob([
        "NCollection/*.cxx",
        "Precision/*.cxx",
        "Standard/*.cxx",
        "TColStd/*.cxx",
        "TCollection/*.cxx",
    ]),
    hdrs = glob([
        "NCollection/*.gxx",
        "NCollection/*.pxx",
        "NCollection/*.hxx",
        "NCollection/*.lxx",
        "NCollection/*.h",
        "Precision/*.gxx",
        "Precision/*.pxx",
        "Precision/*.hxx",
        "Precision/*.lxx",
        "Precision/*.h",
        "Standard/*.gxx",
        "Standard/*.pxx",
        "Standard/*.hxx",
        "Standard/*.lxx",
        "Standard/*.h",
        "TColStd/*.gxx",
        "TColStd/*.pxx",
        "TColStd/*.hxx",
        "TColStd/*.lxx",
        "TColStd/*.h",
        "TCollection/*.gxx",
        "TCollection/*.pxx",
        "TCollection/*.hxx",
        "TCollection/*.lxx",
        "TCollection/*.h",
    ]),
    includes = [
        "NCollection/",
        "Precision/",
        "Standard/",
        "TColStd/",
        "TCollection/",
    ],
    deps = [
    ],
)

cc_library(
    name = "NLPlate",
    srcs = glob([
        "NLPlate/*.cxx",
    ]),
    hdrs = glob([
        "NLPlate/*.gxx",
        "NLPlate/*.pxx",
        "NLPlate/*.hxx",
        "NLPlate/*.lxx",
        "NLPlate/*.h",
    ]),
    includes = [
        "NLPlate/",
    ],
    deps = [
        ":Adaptor2d",
        ":NCollection",
        ":Plate",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "OSD",
    srcs = glob([
        "OSD/*.cxx",
    ]),
    hdrs = glob([
        "OSD/*.gxx",
        "OSD/*.pxx",
        "OSD/*.hxx",
        "OSD/*.lxx",
        "OSD/*.h",
    ]),
    includes = [
        "OSD/",
    ],
    deps = [
        ":NCollection",
        ":Quantity",
    ],
)

cc_library(
    name = "OpenGl",
    srcs = glob([
        "OpenGl/*.cxx",
    ]),
    hdrs = glob([
        "OpenGl/*.gxx",
        "OpenGl/*.pxx",
        "OpenGl/*.hxx",
        "OpenGl/*.lxx",
        "OpenGl/*.h",
    ]),
    includes = [
        "OpenGl/",
    ],
    deps = [
        ":Aspect",
        ":BVH",
        ":Cocoa",
        ":Font",
        ":Image",
        ":InterfaceGraphic",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":Shaders",
        ":WNT",
        ":Xw",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "PCDM",
    srcs = glob([
        "PCDM/*.cxx",
    ]),
    hdrs = glob([
        "PCDM/*.gxx",
        "PCDM/*.pxx",
        "PCDM/*.hxx",
        "PCDM/*.lxx",
        "PCDM/*.h",
    ]),
    includes = [
        "PCDM/",
    ],
    deps = [
        ":CDM",
        ":FSD",
        ":LDOM",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Plugin",
        ":Resource",
        ":Storage",
        ":UTL",
    ],
)

cc_library(
    name = "PLib",
    srcs = glob([
        "PLib/*.cxx",
    ]),
    hdrs = glob([
        "PLib/*.gxx",
        "PLib/*.pxx",
        "PLib/*.hxx",
        "PLib/*.lxx",
        "PLib/*.h",
    ]),
    includes = [
        "PLib/",
    ],
    deps = [
        ":GeomAbs",
        ":NCollection",
        ":TColgp",
        ":math",
    ],
)

cc_library(
    name = "Plate",
    srcs = glob([
        "Plate/*.cxx",
    ]),
    hdrs = glob([
        "Plate/*.gxx",
        "Plate/*.pxx",
        "Plate/*.hxx",
        "Plate/*.lxx",
        "Plate/*.h",
    ]),
    includes = [
        "Plate/",
    ],
    deps = [
        ":NCollection",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Plugin",
    srcs = glob([
        "Plugin/*.cxx",
    ]),
    hdrs = glob([
        "Plugin/*.gxx",
        "Plugin/*.pxx",
        "Plugin/*.hxx",
        "Plugin/*.lxx",
        "Plugin/*.h",
    ]),
    includes = [
        "Plugin/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
        ":Resource",
    ],
)

cc_library(
    name = "Poly",
    srcs = glob([
        "Poly/*.cxx",
    ]),
    hdrs = glob([
        "Poly/*.gxx",
        "Poly/*.pxx",
        "Poly/*.hxx",
        "Poly/*.lxx",
        "Poly/*.h",
    ]),
    includes = [
        "Poly/",
    ],
    deps = [
        ":NCollection",
        ":TColgp",
        ":TShort",
        ":gp",
    ],
)

cc_library(
    name = "ProjLib",
    srcs = glob([
        "ProjLib/*.cxx",
    ]),
    hdrs = glob([
        "ProjLib/*.gxx",
        "ProjLib/*.pxx",
        "ProjLib/*.hxx",
        "ProjLib/*.lxx",
        "ProjLib/*.h",
    ]),
    includes = [
        "ProjLib/",
    ],
    deps = [
        ":Adaptor2d",
        ":AppCont",
        ":AppParCurves",
        ":Approx",
        ":BSplCLib",
        ":Convert",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":GeomLib",
        ":NCollection",
        ":OSD",
        ":PLib",
        ":StdFail",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "QABugs",
    srcs = glob([
        "QABugs/*.cxx",
    ]),
    hdrs = glob([
        "QABugs/*.gxx",
        "QABugs/*.pxx",
        "QABugs/*.hxx",
        "QABugs/*.lxx",
        "QABugs/*.h",
    ]),
    includes = [
        "QABugs/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":AppStd",
        ":ApproxInt",
        ":Aspect",
        ":BOPAlgo",
        ":BRep",
        ":BRepAlgo",
        ":BRepAlgoAPI",
        ":BRepFeat",
        ":BRepFilletAPI",
        ":BRepOffset",
        ":BRepOffsetAPI",
        ":BRepPrimAPI",
        ":Bnd",
        ":CDF",
        ":ChFi3d",
        ":DDF",
        ":DDataStd",
        ":DDocStd",
        ":Draw",
        ":ElSLib",
        ":Expr",
        ":ExprIntrp",
        ":FSD",
        ":Font",
        ":GC",
        ":GCE2d",
        ":GProp",
        ":GccEnt",
        ":Geom2dAPI",
        ":Geom2dGcc",
        ":GeomAbs",
        ":GeomFill",
        ":GeomPlate",
        ":GeomProjLib",
        ":HLRAlgo",
        ":HLRAppli",
        ":IFSelect",
        ":IGESBasic",
        ":IGESControl",
        ":IGESToBRep",
        ":Image",
        ":IntCurveSurface",
        ":IntRes2d",
        ":IntTools",
        ":Interface",
        ":LDOM",
        ":Law",
        ":Message",
        ":NCollection",
        ":OSD",
        ":PCDM",
        ":Plugin",
        ":Poly",
        ":ProjLib",
        ":Resource",
        ":STEPCAFControl",
        ":STEPControl",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeExtend",
        ":ShapeProcess",
        ":ShapeUpgrade",
        ":StdSelect",
        ":TColgp",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TNaming",
        ":TPrsStd",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":Units",
        ":ViewerTest",
        ":XCAFDimTolObjects",
        ":XCAFPrs",
        ":XCAFView",
        ":XSControl",
        ":XSDRAW",
        ":XmlDrivers",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "QADNaming",
    srcs = glob([
        "QADNaming/*.cxx",
    ]),
    hdrs = glob([
        "QADNaming/*.gxx",
        "QADNaming/*.pxx",
        "QADNaming/*.hxx",
        "QADNaming/*.lxx",
        "QADNaming/*.h",
    ]),
    includes = [
        "QADNaming/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":DDF",
        ":DNaming",
        ":Draw",
        ":NCollection",
        ":TDF",
        ":TNaming",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "QADraw",
    srcs = glob([
        "QADraw/*.cxx",
    ]),
    hdrs = glob([
        "QADraw/*.gxx",
        "QADraw/*.pxx",
        "QADraw/*.hxx",
        "QADraw/*.lxx",
        "QADraw/*.h",
    ]),
    includes = [
        "QADraw/",
    ],
    deps = [
        ":AIS",
        ":ApproxInt",
        ":Bnd",
        ":Draw",
        ":NCollection",
        ":OSD",
        ":QABugs",
        ":QADNaming",
        ":QANCollection",
        ":TColgp",
        ":TopTools",
        ":TopoDS",
        ":ViewerTest",
    ],
)

cc_library(
    name = "QANCollection",
    srcs = glob([
        "QANCollection/*.cxx",
    ]),
    hdrs = glob([
        "QANCollection/*.gxx",
        "QANCollection/*.pxx",
        "QANCollection/*.hxx",
        "QANCollection/*.lxx",
        "QANCollection/*.h",
    ]),
    includes = [
        "QANCollection/",
    ],
    deps = [
        ":Adaptor2d",
        ":Draw",
        ":Message",
        ":NCollection",
        ":OSD",
        ":TColgp",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "Quantity",
    srcs = glob([
        "Quantity/*.cxx",
    ]),
    hdrs = glob([
        "Quantity/*.gxx",
        "Quantity/*.pxx",
        "Quantity/*.hxx",
        "Quantity/*.lxx",
        "Quantity/*.h",
    ]),
    includes = [
        "Quantity/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "RWHeaderSection",
    srcs = glob([
        "RWHeaderSection/*.cxx",
    ]),
    hdrs = glob([
        "RWHeaderSection/*.gxx",
        "RWHeaderSection/*.pxx",
        "RWHeaderSection/*.hxx",
        "RWHeaderSection/*.lxx",
        "RWHeaderSection/*.h",
    ]),
    includes = [
        "RWHeaderSection/",
    ],
    deps = [
        ":HeaderSection",
        ":Interface",
        ":NCollection",
        ":StepData",
    ],
)

cc_library(
    name = "RWStepAP203",
    srcs = glob([
        "RWStepAP203/*.cxx",
    ]),
    hdrs = glob([
        "RWStepAP203/*.gxx",
        "RWStepAP203/*.pxx",
        "RWStepAP203/*.hxx",
        "RWStepAP203/*.lxx",
        "RWStepAP203/*.h",
    ]),
    includes = [
        "RWStepAP203/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepAP203",
        ":StepBasic",
        ":StepData",
    ],
)

cc_library(
    name = "RWStepAP214",
    srcs = glob([
        "RWStepAP214/*.cxx",
    ]),
    hdrs = glob([
        "RWStepAP214/*.gxx",
        "RWStepAP214/*.pxx",
        "RWStepAP214/*.hxx",
        "RWStepAP214/*.lxx",
        "RWStepAP214/*.h",
    ]),
    includes = [
        "RWStepAP214/",
    ],
    deps = [
        ":Interface",
        ":MoniTool",
        ":NCollection",
        ":RWHeaderSection",
        ":RWStepAP203",
        ":RWStepAP242",
        ":RWStepBasic",
        ":RWStepDimTol",
        ":RWStepElement",
        ":RWStepFEA",
        ":RWStepGeom",
        ":RWStepRepr",
        ":RWStepShape",
        ":RWStepVisual",
        ":StepAP203",
        ":StepAP214",
        ":StepBasic",
        ":StepData",
        ":StepDimTol",
        ":StepElement",
        ":StepFEA",
        ":StepGeom",
        ":StepVisual",
    ],
)

cc_library(
    name = "RWStepAP242",
    srcs = glob([
        "RWStepAP242/*.cxx",
    ]),
    hdrs = glob([
        "RWStepAP242/*.gxx",
        "RWStepAP242/*.pxx",
        "RWStepAP242/*.hxx",
        "RWStepAP242/*.lxx",
        "RWStepAP242/*.h",
    ]),
    includes = [
        "RWStepAP242/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepAP214",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepBasic",
    srcs = glob([
        "RWStepBasic/*.cxx",
    ]),
    hdrs = glob([
        "RWStepBasic/*.gxx",
        "RWStepBasic/*.pxx",
        "RWStepBasic/*.hxx",
        "RWStepBasic/*.lxx",
        "RWStepBasic/*.h",
    ]),
    includes = [
        "RWStepBasic/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
    ],
)

cc_library(
    name = "RWStepDimTol",
    srcs = glob([
        "RWStepDimTol/*.cxx",
    ]),
    hdrs = glob([
        "RWStepDimTol/*.gxx",
        "RWStepDimTol/*.pxx",
        "RWStepDimTol/*.hxx",
        "RWStepDimTol/*.lxx",
        "RWStepDimTol/*.h",
    ]),
    includes = [
        "RWStepDimTol/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepDimTol",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepElement",
    srcs = glob([
        "RWStepElement/*.cxx",
    ]),
    hdrs = glob([
        "RWStepElement/*.gxx",
        "RWStepElement/*.pxx",
        "RWStepElement/*.hxx",
        "RWStepElement/*.lxx",
        "RWStepElement/*.h",
    ]),
    includes = [
        "RWStepElement/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepData",
        ":StepElement",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepFEA",
    srcs = glob([
        "RWStepFEA/*.cxx",
    ]),
    hdrs = glob([
        "RWStepFEA/*.gxx",
        "RWStepFEA/*.pxx",
        "RWStepFEA/*.hxx",
        "RWStepFEA/*.lxx",
        "RWStepFEA/*.h",
    ]),
    includes = [
        "RWStepFEA/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepElement",
        ":StepFEA",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepGeom",
    srcs = glob([
        "RWStepGeom/*.cxx",
    ]),
    hdrs = glob([
        "RWStepGeom/*.gxx",
        "RWStepGeom/*.pxx",
        "RWStepGeom/*.hxx",
        "RWStepGeom/*.lxx",
        "RWStepGeom/*.h",
    ]),
    includes = [
        "RWStepGeom/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepRepr",
    srcs = glob([
        "RWStepRepr/*.cxx",
    ]),
    hdrs = glob([
        "RWStepRepr/*.gxx",
        "RWStepRepr/*.pxx",
        "RWStepRepr/*.hxx",
        "RWStepRepr/*.lxx",
        "RWStepRepr/*.h",
    ]),
    includes = [
        "RWStepRepr/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepShape",
    srcs = glob([
        "RWStepShape/*.cxx",
    ]),
    hdrs = glob([
        "RWStepShape/*.gxx",
        "RWStepShape/*.pxx",
        "RWStepShape/*.hxx",
        "RWStepShape/*.lxx",
        "RWStepShape/*.h",
    ]),
    includes = [
        "RWStepShape/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "RWStepVisual",
    srcs = glob([
        "RWStepVisual/*.cxx",
    ]),
    hdrs = glob([
        "RWStepVisual/*.gxx",
        "RWStepVisual/*.pxx",
        "RWStepVisual/*.hxx",
        "RWStepVisual/*.lxx",
        "RWStepVisual/*.h",
    ]),
    includes = [
        "RWStepVisual/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
        ":StepVisual",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "RWStl",
    srcs = glob([
        "RWStl/*.cxx",
    ]),
    hdrs = glob([
        "RWStl/*.gxx",
        "RWStl/*.pxx",
        "RWStl/*.hxx",
        "RWStl/*.lxx",
        "RWStl/*.h",
    ]),
    includes = [
        "RWStl/",
    ],
    deps = [
        ":FSD",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":gp",
    ],
)

cc_library(
    name = "Resource",
    srcs = glob([
        "Resource/*.cxx",
    ]),
    hdrs = glob([
        "Resource/*.gxx",
        "Resource/*.pxx",
        "Resource/*.hxx",
        "Resource/*.lxx",
        "Resource/*.h",
    ]),
    includes = [
        "Resource/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "SHMessage",
    srcs = glob([
        "SHMessage/*.cxx",
    ]),
    hdrs = glob([
        "SHMessage/*.gxx",
        "SHMessage/*.pxx",
        "SHMessage/*.hxx",
        "SHMessage/*.lxx",
        "SHMessage/*.h",
    ]),
    includes = [
        "SHMessage/",
    ],
    deps = [
    ],
)

cc_library(
    name = "STEPCAFControl",
    srcs = glob([
        "STEPCAFControl/*.cxx",
    ]),
    hdrs = glob([
        "STEPCAFControl/*.gxx",
        "STEPCAFControl/*.pxx",
        "STEPCAFControl/*.hxx",
        "STEPCAFControl/*.lxx",
        "STEPCAFControl/*.h",
    ]),
    includes = [
        "STEPCAFControl/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Bnd",
        ":GeomToStep",
        ":HeaderSection",
        ":IFSelect",
        ":Interface",
        ":Message",
        ":MoniTool",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":STEPConstruct",
        ":STEPControl",
        ":ShapeAnalysis",
        ":StepAP214",
        ":StepBasic",
        ":StepData",
        ":StepDimTol",
        ":StepGeom",
        ":StepToGeom",
        ":StepVisual",
        ":TColgp",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TNaming",
        ":TPrsStd",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XCAFDimTolObjects",
        ":XCAFPrs",
        ":XCAFView",
        ":XSAlgo",
        ":XSControl",
        ":gp",
    ],
)

cc_library(
    name = "STEPConstruct",
    srcs = glob([
        "STEPConstruct/*.cxx",
    ]),
    hdrs = glob([
        "STEPConstruct/*.gxx",
        "STEPConstruct/*.pxx",
        "STEPConstruct/*.hxx",
        "STEPConstruct/*.lxx",
        "STEPConstruct/*.h",
    ]),
    includes = [
        "STEPConstruct/",
    ],
    deps = [
        ":APIHeaderSection",
        ":GeomToStep",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":StepAP203",
        ":StepAP214",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
        ":StepVisual",
        ":TopLoc",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":UnitsMethods",
        ":XSControl",
        ":gp",
    ],
)

cc_library(
    name = "STEPControl",
    srcs = glob([
        "STEPControl/*.cxx",
        "StepToTopoDS/*.cxx",
    ]),
    hdrs = glob([
        "STEPControl/*.gxx",
        "STEPControl/*.pxx",
        "STEPControl/*.hxx",
        "STEPControl/*.lxx",
        "STEPControl/*.h",
        "StepToTopoDS/*.gxx",
        "StepToTopoDS/*.pxx",
        "StepToTopoDS/*.hxx",
        "StepToTopoDS/*.lxx",
        "StepToTopoDS/*.h",
    ]),
    includes = [
        "STEPControl/",
        "StepToTopoDS/",
    ],
    deps = [
        ":APIHeaderSection",
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomToStep",
        ":HeaderSection",
        ":IFSelect",
        ":Interface",
        ":Message",
        ":MoniTool",
        ":NCollection",
        ":OSD",
        ":RWHeaderSection",
        ":RWStepAP214",
        ":STEPConstruct",
        ":STEPEdit",
        ":STEPSelections",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeExtend",
        ":ShapeProcess",
        ":StdFail",
        ":StepBasic",
        ":StepData",
        ":StepDimTol",
        ":StepGeom",
        ":StepSelect",
        ":StepToGeom",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":TopoDSToStep",
        ":Transfer",
        ":TransferBRep",
        ":UnitsMethods",
        ":XSAlgo",
        ":XSControl",
        ":gp",
    ],
)

cc_library(
    name = "STEPEdit",
    srcs = glob([
        "STEPEdit/*.cxx",
    ]),
    hdrs = glob([
        "STEPEdit/*.gxx",
        "STEPEdit/*.pxx",
        "STEPEdit/*.hxx",
        "STEPEdit/*.lxx",
        "STEPEdit/*.h",
    ]),
    includes = [
        "STEPEdit/",
    ],
    deps = [
        ":APIHeaderSection",
        ":IFSelect",
        ":Interface",
        ":NCollection",
        ":STEPConstruct",
        ":StepAP214",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
        ":StepSelect",
    ],
)

cc_library(
    name = "STEPSelections",
    srcs = glob([
        "STEPSelections/*.cxx",
    ]),
    hdrs = glob([
        "STEPSelections/*.gxx",
        "STEPSelections/*.pxx",
        "STEPSelections/*.hxx",
        "STEPSelections/*.lxx",
        "STEPSelections/*.h",
    ]),
    includes = [
        "STEPSelections/",
    ],
    deps = [
        ":IFSelect",
        ":Interface",
        ":NCollection",
        ":RWStepAP214",
        ":STEPConstruct",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
        ":StepSelect",
        ":Transfer",
        ":XSControl",
    ],
)

cc_library(
    name = "SWDRAW",
    srcs = glob([
        "SWDRAW/*.cxx",
    ]),
    hdrs = glob([
        "SWDRAW/*.gxx",
        "SWDRAW/*.pxx",
        "SWDRAW/*.hxx",
        "SWDRAW/*.lxx",
        "SWDRAW/*.h",
    ]),
    includes = [
        "SWDRAW/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepTest",
        ":Draw",
        ":GProp",
        ":GeomAbs",
        ":GeomLib",
        ":Message",
        ":NCollection",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":ShapeProcess",
        ":ShapeProcessAPI",
        ":ShapeUpgrade",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "Shaders",
    srcs = glob([
        "Shaders/*.cxx",
    ]),
    hdrs = glob([
        "Shaders/*.gxx",
        "Shaders/*.pxx",
        "Shaders/*.hxx",
        "Shaders/*.lxx",
        "Shaders/*.h",
    ]),
    includes = [
        "Shaders/",
    ],
    deps = [
    ],
)

cc_library(
    name = "ShapeAlgo",
    srcs = glob([
        "ShapeAlgo/*.cxx",
    ]),
    hdrs = glob([
        "ShapeAlgo/*.gxx",
        "ShapeAlgo/*.pxx",
        "ShapeAlgo/*.hxx",
        "ShapeAlgo/*.lxx",
        "ShapeAlgo/*.h",
    ]),
    includes = [
        "ShapeAlgo/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":GeomAbs",
        ":IntSurf",
        ":NCollection",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":ShapeUpgrade",
        ":TColgp",
        ":TopExp",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapeAnalysis",
    srcs = glob([
        "ShapeAnalysis/*.cxx",
        "ShapeConstruct/*.cxx",
        "ShapeFix/*.cxx",
    ]),
    hdrs = glob([
        "ShapeAnalysis/*.gxx",
        "ShapeAnalysis/*.pxx",
        "ShapeAnalysis/*.hxx",
        "ShapeAnalysis/*.lxx",
        "ShapeAnalysis/*.h",
        "ShapeConstruct/*.gxx",
        "ShapeConstruct/*.pxx",
        "ShapeConstruct/*.hxx",
        "ShapeConstruct/*.lxx",
        "ShapeConstruct/*.h",
        "ShapeFix/*.gxx",
        "ShapeFix/*.pxx",
        "ShapeFix/*.hxx",
        "ShapeFix/*.lxx",
        "ShapeFix/*.h",
    ]),
    includes = [
        "ShapeAnalysis/",
        "ShapeConstruct/",
        "ShapeFix/",
    ],
    deps = [
        ":Adaptor2d",
        ":Approx",
        ":ApproxInt",
        ":BRep",
        ":BSplCLib",
        ":Bnd",
        ":ElCLib",
        ":ElSLib",
        ":GProp",
        ":Geom2dAPI",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomLib",
        ":GeomProjLib",
        ":IntCurve",
        ":IntRes2d",
        ":Message",
        ":NCollection",
        ":Poly",
        ":ProjLib",
        ":ShapeBuild",
        ":ShapeExtend",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapeBuild",
    srcs = glob([
        "ShapeBuild/*.cxx",
    ]),
    hdrs = glob([
        "ShapeBuild/*.gxx",
        "ShapeBuild/*.pxx",
        "ShapeBuild/*.hxx",
        "ShapeBuild/*.lxx",
        "ShapeBuild/*.h",
    ]),
    includes = [
        "ShapeBuild/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":NCollection",
        ":ShapeExtend",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapeCustom",
    srcs = glob([
        "ShapeCustom/*.cxx",
    ]),
    hdrs = glob([
        "ShapeCustom/*.gxx",
        "ShapeCustom/*.pxx",
        "ShapeCustom/*.hxx",
        "ShapeCustom/*.lxx",
        "ShapeCustom/*.h",
    ]),
    includes = [
        "ShapeCustom/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":ElSLib",
        ":GeomAbs",
        ":Message",
        ":NCollection",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeExtend",
        ":TColgp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapeExtend",
    srcs = glob([
        "ShapeExtend/*.cxx",
    ]),
    hdrs = glob([
        "ShapeExtend/*.gxx",
        "ShapeExtend/*.pxx",
        "ShapeExtend/*.hxx",
        "ShapeExtend/*.lxx",
        "ShapeExtend/*.h",
    ]),
    includes = [
        "ShapeExtend/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":GeomAbs",
        ":Message",
        ":NCollection",
        ":SHMessage",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapePersistent",
    srcs = glob([
        "ShapePersistent/*.cxx",
    ]),
    hdrs = glob([
        "ShapePersistent/*.gxx",
        "ShapePersistent/*.pxx",
        "ShapePersistent/*.hxx",
        "ShapePersistent/*.lxx",
        "ShapePersistent/*.h",
    ]),
    includes = [
        "ShapePersistent/",
    ],
    deps = [
        ":Adaptor2d",
        ":BRep",
        ":NCollection",
        ":Poly",
        ":StdLPersistent",
        ":StdObjMgt",
        ":StdObject",
        ":TColgp",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "ShapeProcess",
    srcs = glob([
        "ShapeProcess/*.cxx",
    ]),
    hdrs = glob([
        "ShapeProcess/*.gxx",
        "ShapeProcess/*.pxx",
        "ShapeProcess/*.hxx",
        "ShapeProcess/*.lxx",
        "ShapeProcess/*.h",
    ]),
    includes = [
        "ShapeProcess/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":GeomAbs",
        ":Message",
        ":NCollection",
        ":Resource",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":ShapeUpgrade",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
    ],
)

cc_library(
    name = "ShapeProcessAPI",
    srcs = glob([
        "ShapeProcessAPI/*.cxx",
    ]),
    hdrs = glob([
        "ShapeProcessAPI/*.gxx",
        "ShapeProcessAPI/*.pxx",
        "ShapeProcessAPI/*.hxx",
        "ShapeProcessAPI/*.lxx",
        "ShapeProcessAPI/*.h",
    ]),
    includes = [
        "ShapeProcessAPI/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":Resource",
        ":ShapeProcess",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
    ],
)

cc_library(
    name = "ShapeUpgrade",
    srcs = glob([
        "ShapeUpgrade/*.cxx",
    ]),
    hdrs = glob([
        "ShapeUpgrade/*.gxx",
        "ShapeUpgrade/*.pxx",
        "ShapeUpgrade/*.hxx",
        "ShapeUpgrade/*.lxx",
        "ShapeUpgrade/*.h",
    ]),
    includes = [
        "ShapeUpgrade/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BSplCLib",
        ":Bnd",
        ":GC",
        ":GProp",
        ":GeomAbs",
        ":GeomLib",
        ":Message",
        ":NCollection",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "StdDrivers",
    srcs = glob([
        "StdDrivers/*.cxx",
    ]),
    hdrs = glob([
        "StdDrivers/*.gxx",
        "StdDrivers/*.pxx",
        "StdDrivers/*.hxx",
        "StdDrivers/*.lxx",
        "StdDrivers/*.h",
    ]),
    includes = [
        "StdDrivers/",
    ],
    deps = [
        ":NCollection",
        ":PCDM",
        ":Plugin",
        ":ShapePersistent",
        ":StdLDrivers",
        ":StdLPersistent",
        ":StdObject",
        ":TDocStd",
    ],
)

cc_library(
    name = "StdFail",
    srcs = glob([
        "StdFail/*.cxx",
    ]),
    hdrs = glob([
        "StdFail/*.gxx",
        "StdFail/*.pxx",
        "StdFail/*.hxx",
        "StdFail/*.lxx",
        "StdFail/*.h",
    ]),
    includes = [
        "StdFail/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "StdLDrivers",
    srcs = glob([
        "StdLDrivers/*.cxx",
    ]),
    hdrs = glob([
        "StdLDrivers/*.gxx",
        "StdLDrivers/*.pxx",
        "StdLDrivers/*.hxx",
        "StdLDrivers/*.lxx",
        "StdLDrivers/*.h",
    ]),
    includes = [
        "StdLDrivers/",
    ],
    deps = [
        ":NCollection",
        ":PCDM",
        ":Plugin",
        ":StdLPersistent",
        ":StdObjMgt",
        ":Storage",
        ":TDocStd",
    ],
)

cc_library(
    name = "StdLPersistent",
    srcs = glob([
        "StdLPersistent/*.cxx",
    ]),
    hdrs = glob([
        "StdLPersistent/*.gxx",
        "StdLPersistent/*.pxx",
        "StdLPersistent/*.hxx",
        "StdLPersistent/*.lxx",
        "StdLPersistent/*.h",
    ]),
    includes = [
        "StdLPersistent/",
    ],
    deps = [
        ":NCollection",
        ":StdObjMgt",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TFunction",
    ],
)

cc_library(
    name = "StdObjMgt",
    srcs = glob([
        "StdObjMgt/*.cxx",
    ]),
    hdrs = glob([
        "StdObjMgt/*.gxx",
        "StdObjMgt/*.pxx",
        "StdObjMgt/*.hxx",
        "StdObjMgt/*.lxx",
        "StdObjMgt/*.h",
    ]),
    includes = [
        "StdObjMgt/",
    ],
    deps = [
        ":NCollection",
        ":Storage",
        ":TDF",
    ],
)

cc_library(
    name = "StdObject",
    srcs = glob([
        "StdObject/*.cxx",
        "StdPersistent/*.cxx",
    ]),
    hdrs = glob([
        "StdObject/*.gxx",
        "StdObject/*.pxx",
        "StdObject/*.hxx",
        "StdObject/*.lxx",
        "StdObject/*.h",
        "StdPersistent/*.gxx",
        "StdPersistent/*.pxx",
        "StdPersistent/*.hxx",
        "StdPersistent/*.lxx",
        "StdPersistent/*.h",
    ]),
    includes = [
        "StdObject/",
        "StdPersistent/",
    ],
    deps = [
        ":NCollection",
        ":StdLPersistent",
        ":StdObjMgt",
        ":TDataStd",
        ":TDataXtd",
        ":TNaming",
        ":TopLoc",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "StdPrs",
    srcs = glob([
        "StdPrs/*.cxx",
    ]),
    hdrs = glob([
        "StdPrs/*.gxx",
        "StdPrs/*.pxx",
        "StdPrs/*.hxx",
        "StdPrs/*.lxx",
        "StdPrs/*.h",
    ]),
    includes = [
        "StdPrs/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":Font",
        ":GeomAbs",
        ":GeomLib",
        ":HLRAlgo",
        ":Hatch",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":TColgp",
        ":TShort",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "StdSelect",
    srcs = glob([
        "StdSelect/*.cxx",
    ]),
    hdrs = glob([
        "StdSelect/*.gxx",
        "StdSelect/*.pxx",
        "StdSelect/*.hxx",
        "StdSelect/*.lxx",
        "StdSelect/*.h",
    ]),
    includes = [
        "StdSelect/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":Font",
        ":GeomAbs",
        ":Message",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":StdPrs",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "StdStorage",
    srcs = glob([
        "StdStorage/*.cxx",
    ]),
    hdrs = glob([
        "StdStorage/*.gxx",
        "StdStorage/*.pxx",
        "StdStorage/*.hxx",
        "StdStorage/*.lxx",
        "StdStorage/*.h",
    ]),
    includes = [
        "StdStorage/",
    ],
    deps = [
        ":NCollection",
        ":PCDM",
        ":StdDrivers",
        ":StdObjMgt",
        ":Storage",
    ],
)

cc_library(
    name = "StepAP203",
    srcs = glob([
        "StepAP203/*.cxx",
    ]),
    hdrs = glob([
        "StepAP203/*.gxx",
        "StepAP203/*.pxx",
        "StepAP203/*.hxx",
        "StepAP203/*.lxx",
        "StepAP203/*.h",
    ]),
    includes = [
        "StepAP203/",
    ],
    deps = [
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "StepAP209",
    srcs = glob([
        "StepAP209/*.cxx",
    ]),
    hdrs = glob([
        "StepAP209/*.gxx",
        "StepAP209/*.pxx",
        "StepAP209/*.hxx",
        "StepAP209/*.lxx",
        "StepAP209/*.h",
    ]),
    includes = [
        "StepAP209/",
    ],
    deps = [
        ":HeaderSection",
        ":Interface",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":STEPConstruct",
        ":StepAP203",
        ":StepAP214",
        ":StepBasic",
        ":StepData",
        ":StepElement",
        ":StepFEA",
        ":StepGeom",
        ":XSControl",
    ],
)

cc_library(
    name = "StepAP214",
    srcs = glob([
        "StepAP214/*.cxx",
        "StepAP242/*.cxx",
    ]),
    hdrs = glob([
        "StepAP214/*.gxx",
        "StepAP214/*.pxx",
        "StepAP214/*.hxx",
        "StepAP214/*.lxx",
        "StepAP214/*.h",
        "StepAP242/*.gxx",
        "StepAP242/*.pxx",
        "StepAP242/*.hxx",
        "StepAP242/*.lxx",
        "StepAP242/*.h",
    ]),
    includes = [
        "StepAP214/",
        "StepAP242/",
    ],
    deps = [
        ":HeaderSection",
        ":Interface",
        ":NCollection",
        ":StepAP203",
        ":StepBasic",
        ":StepData",
        ":StepDimTol",
        ":StepElement",
        ":StepFEA",
        ":StepGeom",
        ":StepVisual",
    ],
)

cc_library(
    name = "StepBasic",
    srcs = glob([
        "StepBasic/*.cxx",
    ]),
    hdrs = glob([
        "StepBasic/*.gxx",
        "StepBasic/*.pxx",
        "StepBasic/*.hxx",
        "StepBasic/*.lxx",
        "StepBasic/*.h",
    ]),
    includes = [
        "StepBasic/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepData",
    ],
)

cc_library(
    name = "StepData",
    srcs = glob([
        "StepData/*.cxx",
    ]),
    hdrs = glob([
        "StepData/*.gxx",
        "StepData/*.pxx",
        "StepData/*.hxx",
        "StepData/*.lxx",
        "StepData/*.h",
    ]),
    includes = [
        "StepData/",
    ],
    deps = [
        ":Interface",
        ":LibCtl",
        ":Message",
        ":NCollection",
    ],
)

cc_library(
    name = "StepDimTol",
    srcs = glob([
        "StepDimTol/*.cxx",
    ]),
    hdrs = glob([
        "StepDimTol/*.gxx",
        "StepDimTol/*.pxx",
        "StepDimTol/*.hxx",
        "StepDimTol/*.lxx",
        "StepDimTol/*.h",
    ]),
    includes = [
        "StepDimTol/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "StepElement",
    srcs = glob([
        "StepElement/*.cxx",
    ]),
    hdrs = glob([
        "StepElement/*.gxx",
        "StepElement/*.pxx",
        "StepElement/*.hxx",
        "StepElement/*.lxx",
        "StepElement/*.h",
    ]),
    includes = [
        "StepElement/",
    ],
    deps = [
        ":NCollection",
        ":StepData",
        ":StepGeom",
    ],
)

cc_library(
    name = "StepFEA",
    srcs = glob([
        "StepFEA/*.cxx",
    ]),
    hdrs = glob([
        "StepFEA/*.gxx",
        "StepFEA/*.pxx",
        "StepFEA/*.hxx",
        "StepFEA/*.lxx",
        "StepFEA/*.h",
    ]),
    includes = [
        "StepFEA/",
    ],
    deps = [
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepElement",
        ":StepGeom",
    ],
)

cc_library(
    name = "StepFile",
    srcs = glob([
        "StepFile/*.cxx",
    ]),
    hdrs = glob([
        "StepFile/*.gxx",
        "StepFile/*.pxx",
        "StepFile/*.hxx",
        "StepFile/*.lxx",
        "StepFile/*.h",
    ]),
    includes = [
        "StepFile/",
    ],
    deps = [
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":StepData",
        ":Transfer",
    ],
)

cc_library(
    name = "StepGeom",
    srcs = glob([
        "StepGeom/*.cxx",
        "StepRepr/*.cxx",
        "StepShape/*.cxx",
    ]),
    hdrs = glob([
        "StepGeom/*.gxx",
        "StepGeom/*.pxx",
        "StepGeom/*.hxx",
        "StepGeom/*.lxx",
        "StepGeom/*.h",
        "StepRepr/*.gxx",
        "StepRepr/*.pxx",
        "StepRepr/*.hxx",
        "StepRepr/*.lxx",
        "StepRepr/*.h",
        "StepShape/*.gxx",
        "StepShape/*.pxx",
        "StepShape/*.hxx",
        "StepShape/*.lxx",
        "StepShape/*.h",
    ]),
    includes = [
        "StepGeom/",
        "StepRepr/",
        "StepShape/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
    ],
)

cc_library(
    name = "StepSelect",
    srcs = glob([
        "StepSelect/*.cxx",
    ]),
    hdrs = glob([
        "StepSelect/*.gxx",
        "StepSelect/*.pxx",
        "StepSelect/*.hxx",
        "StepSelect/*.lxx",
        "StepSelect/*.h",
    ]),
    includes = [
        "StepSelect/",
    ],
    deps = [
        ":IFSelect",
        ":Interface",
        ":Message",
        ":NCollection",
        ":OSD",
        ":StepData",
        ":StepFile",
    ],
)

cc_library(
    name = "StepToGeom",
    srcs = glob([
        "StepToGeom/*.cxx",
    ]),
    hdrs = glob([
        "StepToGeom/*.gxx",
        "StepToGeom/*.pxx",
        "StepToGeom/*.hxx",
        "StepToGeom/*.lxx",
        "StepToGeom/*.h",
    ]),
    includes = [
        "StepToGeom/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":ElCLib",
        ":NCollection",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":StepGeom",
        ":TopoDS",
        ":UnitsMethods",
        ":gp",
    ],
)

cc_library(
    name = "StepVisual",
    srcs = glob([
        "StepVisual/*.cxx",
    ]),
    hdrs = glob([
        "StepVisual/*.gxx",
        "StepVisual/*.pxx",
        "StepVisual/*.hxx",
        "StepVisual/*.lxx",
        "StepVisual/*.h",
    ]),
    includes = [
        "StepVisual/",
    ],
    deps = [
        ":Interface",
        ":NCollection",
        ":StepBasic",
        ":StepData",
        ":StepGeom",
        ":TColgp",
    ],
)

cc_library(
    name = "StlAPI",
    srcs = glob([
        "StlAPI/*.cxx",
    ]),
    hdrs = glob([
        "StlAPI/*.gxx",
        "StlAPI/*.pxx",
        "StlAPI/*.hxx",
        "StlAPI/*.lxx",
        "StlAPI/*.h",
    ]),
    includes = [
        "StlAPI/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":Bnd",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":RWStl",
        ":TopExp",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "Storage",
    srcs = glob([
        "Storage/*.cxx",
    ]),
    hdrs = glob([
        "Storage/*.gxx",
        "Storage/*.pxx",
        "Storage/*.hxx",
        "Storage/*.lxx",
        "Storage/*.h",
    ]),
    includes = [
        "Storage/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
    ],
)

cc_library(
    name = "Sweep",
    srcs = glob([
        "Sweep/*.cxx",
    ]),
    hdrs = glob([
        "Sweep/*.gxx",
        "Sweep/*.pxx",
        "Sweep/*.hxx",
        "Sweep/*.lxx",
        "Sweep/*.h",
    ]),
    includes = [
        "Sweep/",
    ],
    deps = [
        ":NCollection",
        ":TopAbs",
    ],
)

cc_library(
    name = "TColQuantity",
    srcs = glob([
        "TColQuantity/*.cxx",
    ]),
    hdrs = glob([
        "TColQuantity/*.gxx",
        "TColQuantity/*.pxx",
        "TColQuantity/*.hxx",
        "TColQuantity/*.lxx",
        "TColQuantity/*.h",
    ]),
    includes = [
        "TColQuantity/",
    ],
    deps = [
        ":NCollection",
        ":Quantity",
    ],
)

cc_library(
    name = "TColgp",
    srcs = glob([
        "TColgp/*.cxx",
    ]),
    hdrs = glob([
        "TColgp/*.gxx",
        "TColgp/*.pxx",
        "TColgp/*.hxx",
        "TColgp/*.lxx",
        "TColgp/*.h",
    ]),
    includes = [
        "TColgp/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "TDF",
    srcs = glob([
        "TDF/*.cxx",
    ]),
    hdrs = glob([
        "TDF/*.gxx",
        "TDF/*.pxx",
        "TDF/*.hxx",
        "TDF/*.lxx",
        "TDF/*.h",
    ]),
    includes = [
        "TDF/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "TDataStd",
    srcs = glob([
        "TDataStd/*.cxx",
    ]),
    hdrs = glob([
        "TDataStd/*.gxx",
        "TDataStd/*.pxx",
        "TDataStd/*.hxx",
        "TDataStd/*.lxx",
        "TDataStd/*.h",
    ]),
    includes = [
        "TDataStd/",
    ],
    deps = [
        ":NCollection",
        ":TDF",
    ],
)

cc_library(
    name = "TDataXtd",
    srcs = glob([
        "TDataXtd/*.cxx",
    ]),
    hdrs = glob([
        "TDataXtd/*.gxx",
        "TDataXtd/*.pxx",
        "TDataXtd/*.hxx",
        "TDataXtd/*.lxx",
        "TDataXtd/*.h",
    ]),
    includes = [
        "TDataXtd/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":GeomAbs",
        ":GeomLib",
        ":NCollection",
        ":Poly",
        ":Quantity",
        ":TDF",
        ":TDataStd",
        ":TNaming",
        ":TopAbs",
        ":TopLoc",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TDocStd",
    srcs = glob([
        "TDocStd/*.cxx",
    ]),
    hdrs = glob([
        "TDocStd/*.gxx",
        "TDocStd/*.pxx",
        "TDocStd/*.hxx",
        "TDocStd/*.lxx",
        "TDocStd/*.h",
    ]),
    includes = [
        "TDocStd/",
    ],
    deps = [
        ":CDF",
        ":CDM",
        ":Message",
        ":NCollection",
        ":OSD",
        ":PCDM",
        ":Plugin",
        ":Resource",
        ":TDF",
        ":TDataStd",
    ],
)

cc_library(
    name = "TFunction",
    srcs = glob([
        "TFunction/*.cxx",
    ]),
    hdrs = glob([
        "TFunction/*.gxx",
        "TFunction/*.pxx",
        "TFunction/*.hxx",
        "TFunction/*.lxx",
        "TFunction/*.h",
    ]),
    includes = [
        "TFunction/",
    ],
    deps = [
        ":NCollection",
        ":TDF",
        ":TDataStd",
    ],
)

cc_library(
    name = "TNaming",
    srcs = glob([
        "TNaming/*.cxx",
    ]),
    hdrs = glob([
        "TNaming/*.gxx",
        "TNaming/*.pxx",
        "TNaming/*.hxx",
        "TNaming/*.lxx",
        "TNaming/*.h",
    ]),
    includes = [
        "TNaming/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":IntTools",
        ":NCollection",
        ":TDF",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TObj",
    srcs = glob([
        "TObj/*.cxx",
    ]),
    hdrs = glob([
        "TObj/*.gxx",
        "TObj/*.pxx",
        "TObj/*.hxx",
        "TObj/*.lxx",
        "TObj/*.h",
    ]),
    includes = [
        "TObj/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":OSD",
        ":Resource",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":gp",
    ],
)

cc_library(
    name = "TObjDRAW",
    srcs = glob([
        "TObjDRAW/*.cxx",
    ]),
    hdrs = glob([
        "TObjDRAW/*.gxx",
        "TObjDRAW/*.pxx",
        "TObjDRAW/*.hxx",
        "TObjDRAW/*.lxx",
        "TObjDRAW/*.h",
    ]),
    includes = [
        "TObjDRAW/",
    ],
    deps = [
        ":BinTObjDrivers",
        ":DDocStd",
        ":Draw",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TObj",
        ":XmlTObjDrivers",
    ],
)

cc_library(
    name = "TPrsStd",
    srcs = glob([
        "TPrsStd/*.cxx",
    ]),
    hdrs = glob([
        "TPrsStd/*.gxx",
        "TPrsStd/*.pxx",
        "TPrsStd/*.hxx",
        "TPrsStd/*.lxx",
        "TPrsStd/*.h",
    ]),
    includes = [
        "TPrsStd/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Font",
        ":GC",
        ":GeomAbs",
        ":IntAna",
        ":NCollection",
        ":Quantity",
        ":TDF",
        ":TDataStd",
        ":TDataXtd",
        ":TNaming",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Units",
        ":gp",
    ],
)

cc_library(
    name = "TShort",
    srcs = glob([
        "TShort/*.cxx",
    ]),
    hdrs = glob([
        "TShort/*.gxx",
        "TShort/*.pxx",
        "TShort/*.hxx",
        "TShort/*.lxx",
        "TShort/*.h",
    ]),
    includes = [
        "TShort/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "TopAbs",
    srcs = glob([
        "TopAbs/*.cxx",
    ]),
    hdrs = glob([
        "TopAbs/*.gxx",
        "TopAbs/*.pxx",
        "TopAbs/*.hxx",
        "TopAbs/*.lxx",
        "TopAbs/*.h",
    ]),
    includes = [
        "TopAbs/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "TopBas",
    srcs = glob([
        "TopBas/*.cxx",
    ]),
    hdrs = glob([
        "TopBas/*.gxx",
        "TopBas/*.pxx",
        "TopBas/*.hxx",
        "TopBas/*.lxx",
        "TopBas/*.h",
    ]),
    includes = [
        "TopBas/",
    ],
    deps = [
        ":NCollection",
        ":TopAbs",
    ],
)

cc_library(
    name = "TopClass",
    srcs = glob([
        "TopClass/*.cxx",
    ]),
    hdrs = glob([
        "TopClass/*.gxx",
        "TopClass/*.pxx",
        "TopClass/*.hxx",
        "TopClass/*.lxx",
        "TopClass/*.h",
    ]),
    includes = [
        "TopClass/",
    ],
    deps = [
        ":IntCurveSurface",
        ":IntRes2d",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "TopCnx",
    srcs = glob([
        "TopCnx/*.cxx",
    ]),
    hdrs = glob([
        "TopCnx/*.gxx",
        "TopCnx/*.pxx",
        "TopCnx/*.hxx",
        "TopCnx/*.lxx",
        "TopCnx/*.h",
    ]),
    includes = [
        "TopCnx/",
    ],
    deps = [
        ":NCollection",
        ":TopAbs",
        ":TopTrans",
        ":gp",
    ],
)

cc_library(
    name = "TopExp",
    srcs = glob([
        "TopExp/*.cxx",
    ]),
    hdrs = glob([
        "TopExp/*.gxx",
        "TopExp/*.pxx",
        "TopExp/*.hxx",
        "TopExp/*.lxx",
        "TopExp/*.h",
    ]),
    includes = [
        "TopExp/",
    ],
    deps = [
        ":NCollection",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
    ],
)

cc_library(
    name = "TopLoc",
    srcs = glob([
        "TopLoc/*.cxx",
    ]),
    hdrs = glob([
        "TopLoc/*.gxx",
        "TopLoc/*.pxx",
        "TopLoc/*.hxx",
        "TopLoc/*.lxx",
        "TopLoc/*.h",
    ]),
    includes = [
        "TopLoc/",
    ],
    deps = [
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "TopOpeBRep",
    srcs = glob([
        "TopOpeBRep/*.cxx",
    ]),
    hdrs = glob([
        "TopOpeBRep/*.gxx",
        "TopOpeBRep/*.pxx",
        "TopOpeBRep/*.hxx",
        "TopOpeBRep/*.lxx",
        "TopOpeBRep/*.h",
    ]),
    includes = [
        "TopOpeBRep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepIntCurveSurface",
        ":Bnd",
        ":Geom2dInt",
        ":GeomAbs",
        ":GeomProjLib",
        ":IntCurveSurface",
        ":IntRes2d",
        ":IntSurf",
        ":NCollection",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopOpeBRepDS",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TopOpeBRepBuild",
    srcs = glob([
        "TopOpeBRepBuild/*.cxx",
    ]),
    hdrs = glob([
        "TopOpeBRepBuild/*.gxx",
        "TopOpeBRepBuild/*.pxx",
        "TopOpeBRepBuild/*.hxx",
        "TopOpeBRepBuild/*.lxx",
        "TopOpeBRepBuild/*.h",
    ]),
    includes = [
        "TopOpeBRepBuild/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepLProp",
        ":Bnd",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomProjLib",
        ":IntRes2d",
        ":NCollection",
        ":ProjLib",
        ":StdFail",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopOpeBRepDS",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TopOpeBRepDS",
    srcs = glob([
        "TopOpeBRepDS/*.cxx",
    ]),
    hdrs = glob([
        "TopOpeBRepDS/*.gxx",
        "TopOpeBRepDS/*.pxx",
        "TopOpeBRepDS/*.hxx",
        "TopOpeBRepDS/*.lxx",
        "TopOpeBRepDS/*.h",
    ]),
    includes = [
        "TopOpeBRepDS/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepLProp",
        ":ElCLib",
        ":GeomProjLib",
        ":IntSurf",
        ":NCollection",
        ":ProjLib",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopOpeBRepTool",
        ":TopTools",
        ":TopTrans",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TopOpeBRepTool",
    srcs = glob([
        "TopOpeBRepTool/*.cxx",
    ]),
    hdrs = glob([
        "TopOpeBRepTool/*.gxx",
        "TopOpeBRepTool/*.pxx",
        "TopOpeBRepTool/*.hxx",
        "TopOpeBRepTool/*.lxx",
        "TopOpeBRepTool/*.h",
    ]),
    includes = [
        "TopOpeBRepTool/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepLProp",
        ":BSplCLib",
        ":Bnd",
        ":Draw",
        ":ElCLib",
        ":ElSLib",
        ":Geom2dAPI",
        ":GeomAbs",
        ":GeomLib",
        ":NCollection",
        ":OSD",
        ":ProjLib",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TopTools",
    srcs = glob([
        "TopTools/*.cxx",
    ]),
    hdrs = glob([
        "TopTools/*.gxx",
        "TopTools/*.pxx",
        "TopTools/*.hxx",
        "TopTools/*.lxx",
        "TopTools/*.h",
    ]),
    includes = [
        "TopTools/",
    ],
    deps = [
        ":Adaptor2d",
        ":Bnd",
        ":Message",
        ":NCollection",
        ":TopAbs",
        ":TopLoc",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "TopTrans",
    srcs = glob([
        "TopTrans/*.cxx",
    ]),
    hdrs = glob([
        "TopTrans/*.gxx",
        "TopTrans/*.pxx",
        "TopTrans/*.hxx",
        "TopTrans/*.lxx",
        "TopTrans/*.h",
    ]),
    includes = [
        "TopTrans/",
    ],
    deps = [
        ":NCollection",
        ":TopAbs",
        ":gp",
    ],
)

cc_library(
    name = "TopoDS",
    srcs = glob([
        "TopoDS/*.cxx",
    ]),
    hdrs = glob([
        "TopoDS/*.gxx",
        "TopoDS/*.pxx",
        "TopoDS/*.hxx",
        "TopoDS/*.lxx",
        "TopoDS/*.h",
    ]),
    includes = [
        "TopoDS/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":TopAbs",
        ":TopLoc",
    ],
)

cc_library(
    name = "TopoDSToStep",
    srcs = glob([
        "TopoDSToStep/*.cxx",
    ]),
    hdrs = glob([
        "TopoDSToStep/*.gxx",
        "TopoDSToStep/*.pxx",
        "TopoDSToStep/*.hxx",
        "TopoDSToStep/*.lxx",
        "TopoDSToStep/*.h",
    ]),
    includes = [
        "TopoDSToStep/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Bnd",
        ":GeomToStep",
        ":Interface",
        ":Message",
        ":MoniTool",
        ":NCollection",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeExtend",
        ":StdFail",
        ":StepGeom",
        ":TColgp",
        ":TopExp",
        ":TopLoc",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":UnitsMethods",
        ":gp",
    ],
)

cc_library(
    name = "Transfer",
    srcs = glob([
        "Transfer/*.cxx",
    ]),
    hdrs = glob([
        "Transfer/*.gxx",
        "Transfer/*.pxx",
        "Transfer/*.hxx",
        "Transfer/*.lxx",
        "Transfer/*.h",
    ]),
    includes = [
        "Transfer/",
    ],
    deps = [
        ":Adaptor2d",
        ":Interface",
        ":Message",
        ":NCollection",
    ],
)

cc_library(
    name = "TransferBRep",
    srcs = glob([
        "TransferBRep/*.cxx",
    ]),
    hdrs = glob([
        "TransferBRep/*.gxx",
        "TransferBRep/*.pxx",
        "TransferBRep/*.hxx",
        "TransferBRep/*.lxx",
        "TransferBRep/*.h",
    ]),
    includes = [
        "TransferBRep/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
    ],
)

cc_library(
    name = "UTL",
    srcs = glob([
        "UTL/*.cxx",
    ]),
    hdrs = glob([
        "UTL/*.gxx",
        "UTL/*.pxx",
        "UTL/*.hxx",
        "UTL/*.lxx",
        "UTL/*.h",
    ]),
    includes = [
        "UTL/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
        ":Resource",
        ":Storage",
    ],
)

cc_library(
    name = "Units",
    srcs = glob([
        "Units/*.cxx",
        "UnitsAPI/*.cxx",
    ]),
    hdrs = glob([
        "Units/*.gxx",
        "Units/*.pxx",
        "Units/*.hxx",
        "Units/*.lxx",
        "Units/*.h",
        "UnitsAPI/*.gxx",
        "UnitsAPI/*.pxx",
        "UnitsAPI/*.hxx",
        "UnitsAPI/*.lxx",
        "UnitsAPI/*.h",
    ]),
    includes = [
        "Units/",
        "UnitsAPI/",
    ],
    deps = [
        ":NCollection",
        ":OSD",
        ":Resource",
    ],
)

cc_library(
    name = "UnitsMethods",
    srcs = glob([
        "UnitsMethods/*.cxx",
    ]),
    hdrs = glob([
        "UnitsMethods/*.gxx",
        "UnitsMethods/*.pxx",
        "UnitsMethods/*.hxx",
        "UnitsMethods/*.lxx",
        "UnitsMethods/*.h",
    ]),
    includes = [
        "UnitsMethods/",
    ],
    deps = [
        ":Adaptor2d",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "ViewerTest",
    srcs = glob([
        "ViewerTest/*.cxx",
    ]),
    hdrs = glob([
        "ViewerTest/*.gxx",
        "ViewerTest/*.pxx",
        "ViewerTest/*.hxx",
        "ViewerTest/*.lxx",
        "ViewerTest/*.h",
    ]),
    includes = [
        "ViewerTest/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":BRepAlgo",
        ":BRepFilletAPI",
        ":BRepOffsetAPI",
        ":BRepPrimAPI",
        ":BRepTest",
        ":BiTgte",
        ":ChFi3d",
        ":Cocoa",
        ":Draw",
        ":ElSLib",
        ":FilletSurf",
        ":Font",
        ":GC",
        ":HLRAlgo",
        ":Image",
        ":IntAna",
        ":Message",
        ":NCollection",
        ":OSD",
        ":OpenGl",
        ":Poly",
        ":Quantity",
        ":StdFail",
        ":StdPrs",
        ":StdSelect",
        ":TColgp",
        ":TShort",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":WNT",
        ":Xw",
        ":gp",
    ],
)

cc_library(
    name = "Vrml",
    srcs = glob([
        "Vrml/*.cxx",
    ]),
    hdrs = glob([
        "Vrml/*.gxx",
        "Vrml/*.pxx",
        "Vrml/*.hxx",
        "Vrml/*.lxx",
        "Vrml/*.h",
    ]),
    includes = [
        "Vrml/",
    ],
    deps = [
        ":NCollection",
        ":Quantity",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "VrmlAPI",
    srcs = glob([
        "VrmlAPI/*.cxx",
    ]),
    hdrs = glob([
        "VrmlAPI/*.gxx",
        "VrmlAPI/*.pxx",
        "VrmlAPI/*.hxx",
        "VrmlAPI/*.lxx",
        "VrmlAPI/*.h",
    ]),
    includes = [
        "VrmlAPI/",
    ],
    deps = [
        ":BRep",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Vrml",
        ":VrmlConverter",
        ":VrmlData",
    ],
)

cc_library(
    name = "VrmlConverter",
    srcs = glob([
        "VrmlConverter/*.cxx",
    ]),
    hdrs = glob([
        "VrmlConverter/*.gxx",
        "VrmlConverter/*.pxx",
        "VrmlConverter/*.hxx",
        "VrmlConverter/*.lxx",
        "VrmlConverter/*.h",
    ]),
    includes = [
        "VrmlConverter/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":CSLib",
        ":Font",
        ":HLRAlgo",
        ":Hatch",
        ":NCollection",
        ":Poly",
        ":StdPrs",
        ":TColgp",
        ":TopAbs",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":Vrml",
        ":gp",
        ":math",
    ],
)

cc_library(
    name = "VrmlData",
    srcs = glob([
        "VrmlData/*.cxx",
    ]),
    hdrs = glob([
        "VrmlData/*.gxx",
        "VrmlData/*.pxx",
        "VrmlData/*.hxx",
        "VrmlData/*.lxx",
        "VrmlData/*.h",
    ]),
    includes = [
        "VrmlData/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":BRepPrim",
        ":BRepPrimAPI",
        ":Bnd",
        ":GeomLib",
        ":NCollection",
        ":Poly",
        ":Quantity",
        ":TColgp",
        ":TShort",
        ":TopExp",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "WNT",
    srcs = glob([
        "WNT/*.cxx",
    ]),
    hdrs = glob([
        "WNT/*.gxx",
        "WNT/*.pxx",
        "WNT/*.hxx",
        "WNT/*.lxx",
        "WNT/*.h",
    ]),
    includes = [
        "WNT/",
    ],
    deps = [
        ":Aspect",
        ":NCollection",
        ":Quantity",
    ],
)

cc_library(
    name = "XBRepMesh",
    srcs = glob([
        "XBRepMesh/*.cxx",
    ]),
    hdrs = glob([
        "XBRepMesh/*.gxx",
        "XBRepMesh/*.pxx",
        "XBRepMesh/*.hxx",
        "XBRepMesh/*.lxx",
        "XBRepMesh/*.h",
    ]),
    includes = [
        "XBRepMesh/",
    ],
    deps = [
        ":ApproxInt",
        ":NCollection",
    ],
)

cc_library(
    name = "XCAFApp",
    srcs = glob([
        "XCAFApp/*.cxx",
    ]),
    hdrs = glob([
        "XCAFApp/*.gxx",
        "XCAFApp/*.pxx",
        "XCAFApp/*.hxx",
        "XCAFApp/*.lxx",
        "XCAFApp/*.h",
    ]),
    includes = [
        "XCAFApp/",
    ],
    deps = [
        ":NCollection",
        ":Resource",
        ":TDF",
        ":TDocStd",
        ":TPrsStd",
        ":XCAFDimTolObjects",
        ":XCAFPrs",
    ],
)

cc_library(
    name = "XCAFDimTolObjects",
    srcs = glob([
        "XCAFDimTolObjects/*.cxx",
        "XCAFDoc/*.cxx",
    ]),
    hdrs = glob([
        "XCAFDimTolObjects/*.gxx",
        "XCAFDimTolObjects/*.pxx",
        "XCAFDimTolObjects/*.hxx",
        "XCAFDimTolObjects/*.lxx",
        "XCAFDimTolObjects/*.h",
        "XCAFDoc/*.gxx",
        "XCAFDoc/*.pxx",
        "XCAFDoc/*.hxx",
        "XCAFDoc/*.lxx",
        "XCAFDoc/*.h",
    ]),
    includes = [
        "XCAFDimTolObjects/",
        "XCAFDoc/",
    ],
    deps = [
        ":BRep",
        ":NCollection",
        ":OSD",
        ":Quantity",
        ":TColgp",
        ":TDF",
        ":TDataStd",
        ":TDataXtd",
        ":TDocStd",
        ":TNaming",
        ":TopExp",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":XCAFNoteObjects",
        ":XCAFView",
        ":gp",
    ],
)

cc_library(
    name = "XCAFNoteObjects",
    srcs = glob([
        "XCAFNoteObjects/*.cxx",
    ]),
    hdrs = glob([
        "XCAFNoteObjects/*.gxx",
        "XCAFNoteObjects/*.pxx",
        "XCAFNoteObjects/*.hxx",
        "XCAFNoteObjects/*.lxx",
        "XCAFNoteObjects/*.h",
    ]),
    includes = [
        "XCAFNoteObjects/",
    ],
    deps = [
        ":NCollection",
        ":TopoDS",
        ":gp",
    ],
)

cc_library(
    name = "XCAFPrs",
    srcs = glob([
        "XCAFPrs/*.cxx",
    ]),
    hdrs = glob([
        "XCAFPrs/*.gxx",
        "XCAFPrs/*.pxx",
        "XCAFPrs/*.hxx",
        "XCAFPrs/*.lxx",
        "XCAFPrs/*.h",
    ]),
    includes = [
        "XCAFPrs/",
    ],
    deps = [
        ":AIS",
        ":ApproxInt",
        ":BRep",
        ":Font",
        ":NCollection",
        ":Quantity",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TPrsStd",
        ":TopLoc",
        ":TopTools",
        ":TopoDS",
        ":XCAFDimTolObjects",
        ":gp",
    ],
)

cc_library(
    name = "XCAFView",
    srcs = glob([
        "XCAFView/*.cxx",
    ]),
    hdrs = glob([
        "XCAFView/*.gxx",
        "XCAFView/*.pxx",
        "XCAFView/*.hxx",
        "XCAFView/*.lxx",
        "XCAFView/*.h",
    ]),
    includes = [
        "XCAFView/",
    ],
    deps = [
        ":NCollection",
        ":TColgp",
        ":gp",
    ],
)

cc_library(
    name = "XDEDRAW",
    srcs = glob([
        "XDEDRAW/*.cxx",
    ]),
    hdrs = glob([
        "XDEDRAW/*.gxx",
        "XDEDRAW/*.pxx",
        "XDEDRAW/*.hxx",
        "XDEDRAW/*.lxx",
        "XDEDRAW/*.h",
    ]),
    includes = [
        "XDEDRAW/",
    ],
    deps = [
        ":AIS",
        ":Adaptor2d",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":BinXCAFDrivers",
        ":DDF",
        ":DDocStd",
        ":Draw",
        ":Font",
        ":GProp",
        ":IFSelect",
        ":IGESCAFControl",
        ":IGESControl",
        ":Interface",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":STEPCAFControl",
        ":STEPControl",
        ":TColgp",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TNaming",
        ":TPrsStd",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":ViewerTest",
        ":XCAFDimTolObjects",
        ":XCAFPrs",
        ":XCAFView",
        ":XSControl",
        ":XSDRAW",
        ":XSDRAWIGES",
        ":XSDRAWSTEP",
        ":XmlXCAFDrivers",
        ":gp",
    ],
)

cc_library(
    name = "XSAlgo",
    srcs = glob([
        "XSAlgo/*.cxx",
    ]),
    hdrs = glob([
        "XSAlgo/*.gxx",
        "XSAlgo/*.pxx",
        "XSAlgo/*.hxx",
        "XSAlgo/*.lxx",
        "XSAlgo/*.h",
    ]),
    includes = [
        "XSAlgo/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":Interface",
        ":Message",
        ":NCollection",
        ":Resource",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeBuild",
        ":ShapeCustom",
        ":ShapeExtend",
        ":ShapeProcess",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":UnitsMethods",
        ":gp",
    ],
)

cc_library(
    name = "XSControl",
    srcs = glob([
        "XSControl/*.cxx",
    ]),
    hdrs = glob([
        "XSControl/*.gxx",
        "XSControl/*.pxx",
        "XSControl/*.hxx",
        "XSControl/*.lxx",
        "XSControl/*.h",
    ]),
    includes = [
        "XSControl/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":BRep",
        ":IFSelect",
        ":Interface",
        ":Message",
        ":NCollection",
        ":ShapeAlgo",
        ":ShapeAnalysis",
        ":ShapeCustom",
        ":ShapeExtend",
        ":TopAbs",
        ":TopExp",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":gp",
    ],
)

cc_library(
    name = "XSDRAW",
    srcs = glob([
        "XSDRAW/*.cxx",
    ]),
    hdrs = glob([
        "XSDRAW/*.gxx",
        "XSDRAW/*.pxx",
        "XSDRAW/*.hxx",
        "XSDRAW/*.lxx",
        "XSDRAW/*.h",
    ]),
    includes = [
        "XSDRAW/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Draw",
        ":IFSelect",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TopTools",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XSControl",
        ":gp",
    ],
)

cc_library(
    name = "XSDRAWIGES",
    srcs = glob([
        "XSDRAWIGES/*.cxx",
    ]),
    hdrs = glob([
        "XSDRAWIGES/*.gxx",
        "XSDRAWIGES/*.pxx",
        "XSDRAWIGES/*.hxx",
        "XSDRAWIGES/*.lxx",
        "XSDRAWIGES/*.h",
    ]),
    includes = [
        "XSDRAWIGES/",
    ],
    deps = [
        ":Adaptor2d",
        ":ApproxInt",
        ":Draw",
        ":IFSelect",
        ":IGESBasic",
        ":IGESControl",
        ":IGESSelect",
        ":IGESToBRep",
        ":Interface",
        ":Message",
        ":NCollection",
        ":TopExp",
        ":TopoDS",
        ":Transfer",
        ":XSControl",
        ":XSDRAW",
    ],
)

cc_library(
    name = "XSDRAWSTEP",
    srcs = glob([
        "XSDRAWSTEP/*.cxx",
    ]),
    hdrs = glob([
        "XSDRAWSTEP/*.gxx",
        "XSDRAWSTEP/*.pxx",
        "XSDRAWSTEP/*.hxx",
        "XSDRAWSTEP/*.lxx",
        "XSDRAWSTEP/*.h",
    ]),
    includes = [
        "XSDRAWSTEP/",
    ],
    deps = [
        ":ApproxInt",
        ":Draw",
        ":IFSelect",
        ":Interface",
        ":Message",
        ":NCollection",
        ":STEPControl",
        ":STEPSelections",
        ":StepData",
        ":StepGeom",
        ":StepSelect",
        ":TopExp",
        ":TopoDS",
        ":Transfer",
        ":TransferBRep",
        ":XSControl",
        ":XSDRAW",
    ],
)

cc_library(
    name = "XSDRAWSTLVRML",
    srcs = glob([
        "XSDRAWSTLVRML/*.cxx",
    ]),
    hdrs = glob([
        "XSDRAWSTLVRML/*.gxx",
        "XSDRAWSTLVRML/*.pxx",
        "XSDRAWSTLVRML/*.hxx",
        "XSDRAWSTLVRML/*.lxx",
        "XSDRAWSTLVRML/*.h",
    ]),
    includes = [
        "XSDRAWSTLVRML/",
    ],
    deps = [
        ":AIS",
        ":ApproxInt",
        ":Aspect",
        ":BRep",
        ":Bnd",
        ":Draw",
        ":Font",
        ":MeshVS",
        ":NCollection",
        ":OSD",
        ":Poly",
        ":Quantity",
        ":RWStl",
        ":StdSelect",
        ":StlAPI",
        ":TColgp",
        ":TopoDS",
        ":ViewerTest",
        ":VrmlAPI",
        ":VrmlData",
        ":XSDRAW",
        ":XSDRAWIGES",
        ":XSDRAWSTEP",
    ],
)

cc_library(
    name = "XSMessage",
    srcs = glob([
        "XSMessage/*.cxx",
    ]),
    hdrs = glob([
        "XSMessage/*.gxx",
        "XSMessage/*.pxx",
        "XSMessage/*.hxx",
        "XSMessage/*.lxx",
        "XSMessage/*.h",
    ]),
    includes = [
        "XSMessage/",
    ],
    deps = [
    ],
)

cc_library(
    name = "XmlDrivers",
    srcs = glob([
        "XmlDrivers/*.cxx",
    ]),
    hdrs = glob([
        "XmlDrivers/*.gxx",
        "XmlDrivers/*.pxx",
        "XmlDrivers/*.hxx",
        "XmlDrivers/*.lxx",
        "XmlDrivers/*.h",
    ]),
    includes = [
        "XmlDrivers/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":Plugin",
        ":TDocStd",
        ":TNaming",
        ":XmlLDrivers",
        ":XmlMDataXtd",
        ":XmlMNaming",
        ":XmlObjMgt",
    ],
)

cc_library(
    name = "XmlLDrivers",
    srcs = glob([
        "XmlLDrivers/*.cxx",
        "XmlMDF/*.cxx",
        "XmlMDataStd/*.cxx",
        "XmlMDocStd/*.cxx",
        "XmlMFunction/*.cxx",
    ]),
    hdrs = glob([
        "XmlLDrivers/*.gxx",
        "XmlLDrivers/*.pxx",
        "XmlLDrivers/*.hxx",
        "XmlLDrivers/*.lxx",
        "XmlLDrivers/*.h",
        "XmlMDF/*.gxx",
        "XmlMDF/*.pxx",
        "XmlMDF/*.hxx",
        "XmlMDF/*.lxx",
        "XmlMDF/*.h",
        "XmlMDataStd/*.gxx",
        "XmlMDataStd/*.pxx",
        "XmlMDataStd/*.hxx",
        "XmlMDataStd/*.lxx",
        "XmlMDataStd/*.h",
        "XmlMDocStd/*.gxx",
        "XmlMDocStd/*.pxx",
        "XmlMDocStd/*.hxx",
        "XmlMDocStd/*.lxx",
        "XmlMDocStd/*.h",
        "XmlMFunction/*.gxx",
        "XmlMFunction/*.pxx",
        "XmlMFunction/*.hxx",
        "XmlMFunction/*.lxx",
        "XmlMFunction/*.h",
    ]),
    includes = [
        "XmlLDrivers/",
        "XmlMDF/",
        "XmlMDataStd/",
        "XmlMDocStd/",
        "XmlMFunction/",
    ],
    deps = [
        ":CDM",
        ":LDOM",
        ":Message",
        ":NCollection",
        ":OSD",
        ":PCDM",
        ":Plugin",
        ":Storage",
        ":TDF",
        ":TDataStd",
        ":TDocStd",
        ":TFunction",
        ":UTL",
        ":XmlObjMgt",
    ],
)

cc_library(
    name = "XmlMDataXtd",
    srcs = glob([
        "XmlMDataXtd/*.cxx",
    ]),
    hdrs = glob([
        "XmlMDataXtd/*.gxx",
        "XmlMDataXtd/*.pxx",
        "XmlMDataXtd/*.hxx",
        "XmlMDataXtd/*.lxx",
        "XmlMDataXtd/*.h",
    ]),
    includes = [
        "XmlMDataXtd/",
    ],
    deps = [
        ":LDOM",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TDataStd",
        ":TDataXtd",
        ":TNaming",
        ":XmlLDrivers",
        ":XmlObjMgt",
        ":gp",
    ],
)

cc_library(
    name = "XmlMNaming",
    srcs = glob([
        "XmlMNaming/*.cxx",
    ]),
    hdrs = glob([
        "XmlMNaming/*.gxx",
        "XmlMNaming/*.pxx",
        "XmlMNaming/*.hxx",
        "XmlMNaming/*.lxx",
        "XmlMNaming/*.h",
    ]),
    includes = [
        "XmlMNaming/",
    ],
    deps = [
        ":ApproxInt",
        ":BRep",
        ":LDOM",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TNaming",
        ":TopAbs",
        ":TopTools",
        ":TopoDS",
        ":XmlLDrivers",
        ":XmlObjMgt",
        ":gp",
    ],
)

cc_library(
    name = "XmlMXCAFDoc",
    srcs = glob([
        "XmlMXCAFDoc/*.cxx",
    ]),
    hdrs = glob([
        "XmlMXCAFDoc/*.gxx",
        "XmlMXCAFDoc/*.pxx",
        "XmlMXCAFDoc/*.hxx",
        "XmlMXCAFDoc/*.lxx",
        "XmlMXCAFDoc/*.h",
    ]),
    includes = [
        "XmlMXCAFDoc/",
    ],
    deps = [
        ":LDOM",
        ":Message",
        ":NCollection",
        ":TDF",
        ":TNaming",
        ":TopLoc",
        ":TopTools",
        ":XCAFDimTolObjects",
        ":XmlLDrivers",
        ":XmlMNaming",
        ":XmlObjMgt",
        ":gp",
    ],
)

cc_library(
    name = "XmlObjMgt",
    srcs = glob([
        "XmlObjMgt/*.cxx",
    ]),
    hdrs = glob([
        "XmlObjMgt/*.gxx",
        "XmlObjMgt/*.pxx",
        "XmlObjMgt/*.hxx",
        "XmlObjMgt/*.lxx",
        "XmlObjMgt/*.h",
    ]),
    includes = [
        "XmlObjMgt/",
    ],
    deps = [
        ":LDOM",
        ":NCollection",
        ":gp",
    ],
)

cc_library(
    name = "XmlTObjDrivers",
    srcs = glob([
        "XmlTObjDrivers/*.cxx",
    ]),
    hdrs = glob([
        "XmlTObjDrivers/*.gxx",
        "XmlTObjDrivers/*.pxx",
        "XmlTObjDrivers/*.hxx",
        "XmlTObjDrivers/*.lxx",
        "XmlTObjDrivers/*.h",
    ]),
    includes = [
        "XmlTObjDrivers/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":Plugin",
        ":TDF",
        ":TDocStd",
        ":TObj",
        ":XmlLDrivers",
        ":XmlObjMgt",
    ],
)

cc_library(
    name = "XmlXCAFDrivers",
    srcs = glob([
        "XmlXCAFDrivers/*.cxx",
    ]),
    hdrs = glob([
        "XmlXCAFDrivers/*.gxx",
        "XmlXCAFDrivers/*.pxx",
        "XmlXCAFDrivers/*.hxx",
        "XmlXCAFDrivers/*.lxx",
        "XmlXCAFDrivers/*.h",
    ]),
    includes = [
        "XmlXCAFDrivers/",
    ],
    deps = [
        ":Message",
        ":NCollection",
        ":Plugin",
        ":TDocStd",
        ":XmlDrivers",
        ":XmlLDrivers",
        ":XmlMXCAFDoc",
    ],
)

cc_library(
    name = "Xw",
    srcs = glob([
        "Xw/*.cxx",
    ]),
    hdrs = glob([
        "Xw/*.gxx",
        "Xw/*.pxx",
        "Xw/*.hxx",
        "Xw/*.lxx",
        "Xw/*.h",
    ]),
    includes = [
        "Xw/",
    ],
    deps = [
        ":Aspect",
        ":Message",
        ":NCollection",
        ":Quantity",
    ],
)

cc_library(
    name = "gp",
    srcs = glob([
        "gp/*.cxx",
    ]),
    hdrs = glob([
        "gp/*.gxx",
        "gp/*.pxx",
        "gp/*.hxx",
        "gp/*.lxx",
        "gp/*.h",
    ]),
    includes = [
        "gp/",
    ],
    deps = [
        ":NCollection",
    ],
)

cc_library(
    name = "math",
    srcs = glob([
        "math/*.cxx",
    ]),
    hdrs = glob([
        "math/*.gxx",
        "math/*.pxx",
        "math/*.hxx",
        "math/*.lxx",
        "math/*.h",
    ]),
    includes = [
        "math/",
    ],
    deps = [
        ":NCollection",
        ":StdFail",
        ":gp",
    ],
)

