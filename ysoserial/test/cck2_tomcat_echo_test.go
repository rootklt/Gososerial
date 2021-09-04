package test

import (
	"encoding/hex"
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
	"strings"
	"testing"
)

func TestCCK2TomcatEcho(t *testing.T) {
	ser := gososerial.GetCCK2TomcatEcho("Testecho", "Testcmd")
	serStr := strings.ToUpper(hex.EncodeToString(ser))
	stdStr := "ACED0005737200116A6176612E7574696C2E486173684D61700507DAC1C31660D103000246000A6C6F6164466163746F724900097468726573686F6C6478703F4000000000000C77080000001000000001737200356F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E6B657976616C75652E546965644D6170456E7472798AADD29B39C11FDB0200024C00036B65797400124C6A6176612F6C616E672F4F626A6563743B4C00036D617074000F4C6A6176612F7574696C2F4D61703B78707372003A636F6D2E73756E2E6F72672E6170616368652E78616C616E2E696E7465726E616C2E78736C74632E747261782E54656D706C61746573496D706C09574FC16EACAB3303000649000D5F696E64656E744E756D62657249000E5F7472616E736C6574496E6465785B000A5F62797465636F6465737400035B5B425B00065F636C6173737400125B4C6A6176612F6C616E672F436C6173733B4C00055F6E616D657400124C6A6176612F6C616E672F537472696E673B4C00115F6F757470757450726F706572746965737400164C6A6176612F7574696C2F50726F706572746965733B787000000000FFFFFFFF757200035B5B424BFD19156767DB37020000787000000001757200025B42ACF317F8060854E0020000787000000F94CAFEBABE0000003300EB01001D79736F73657269616C2F50776E657239353330303933333431363030300700010100106A6176612F6C616E672F4F626A65637407000301000A536F7572636546696C6501001850776E657239353330303933333431363030302E6A6176610100097772697465426F6479010017284C6A6176612F6C616E672F4F626A6563743B5B4229560100246F72672E6170616368652E746F6D6361742E7574696C2E6275662E427974654368756E6B08000901000F6A6176612F6C616E672F436C61737307000B010007666F724E616D65010025284C6A6176612F6C616E672F537472696E673B294C6A6176612F6C616E672F436C6173733B0C000D000E0A000C000F01000B6E6577496E7374616E636501001428294C6A6176612F6C616E672F4F626A6563743B0C001100120A000C001301000873657442797465730800150100025B420700170100116A6176612F6C616E672F496E7465676572070019010004545950450100114C6A6176612F6C616E672F436C6173733B0C001B001C09001A001D0100116765744465636C617265644D6574686F64010040284C6A6176612F6C616E672F537472696E673B5B4C6A6176612F6C616E672F436C6173733B294C6A6176612F6C616E672F7265666C6563742F4D6574686F643B0C001F00200A000C00210100063C696E69743E010004284929560C002300240A001A00250100186A6176612F6C616E672F7265666C6563742F4D6574686F64070027010006696E766F6B65010039284C6A6176612F6C616E672F4F626A6563743B5B4C6A6176612F6C616E672F4F626A6563743B294C6A6176612F6C616E672F4F626A6563743B0C0029002A0A0028002B010008676574436C61737301001328294C6A6176612F6C616E672F436C6173733B0C002D002E0A0004002F010007646F57726974650800310100096765744D6574686F640C003300200A000C00340100206A6176612F6C616E672F436C6173734E6F74466F756E64457863657074696F6E0700360100136A6176612E6E696F2E427974654275666665720800380100047772617008003A01001F6A6176612F6C616E672F4E6F537563684D6574686F64457863657074696F6E07003C010004436F646501000A457863657074696F6E730100136A6176612F6C616E672F457863657074696F6E07004001000D537461636B4D61705461626C650100056765744656010038284C6A6176612F6C616E672F4F626A6563743B4C6A6176612F6C616E672F537472696E673B294C6A6176612F6C616E672F4F626A6563743B0100106765744465636C617265644669656C6401002D284C6A6176612F6C616E672F537472696E673B294C6A6176612F6C616E672F7265666C6563742F4669656C643B0C004500460A000C004701001E6A6176612F6C616E672F4E6F537563684669656C64457863657074696F6E07004901000D6765745375706572636C6173730C004B002E0A000C004C010015284C6A6176612F6C616E672F537472696E673B29560C0023004E0A004A004F0100226A6176612F6C616E672F7265666C6563742F41636365737369626C654F626A65637407005101000D73657441636365737369626C65010004285A29560C005300540A005200550100176A6176612F6C616E672F7265666C6563742F4669656C64070057010003676574010026284C6A6176612F6C616E672F4F626A6563743B294C6A6176612F6C616E672F4F626A6563743B0C0059005A0A0058005B0100106A6176612F6C616E672F537472696E6707005D0100032829560C0023005F0A000400600100106A6176612F6C616E672F54687265616407006201000D63757272656E7454687265616401001428294C6A6176612F6C616E672F5468726561643B0C006400650A0063006601000E67657454687265616447726F757001001928294C6A6176612F6C616E672F54687265616447726F75703B0C006800690A0063006A0100077468726561647308006C0C004300440A0002006E0100135B4C6A6176612F6C616E672F5468726561643B0700700100076765744E616D6501001428294C6A6176612F6C616E672F537472696E673B0C007200730A0063007401000465786563080076010008636F6E7461696E7301001B284C6A6176612F6C616E672F4368617253657175656E63653B295A0C007800790A005E007A0100046874747008007C01000674617267657408007E0100126A6176612F6C616E672F52756E6E61626C6507008001000674686973243008008201000768616E646C6572080084010006676C6F62616C08008601000A70726F636573736F727308008801000E6A6176612F7574696C2F4C69737407008A01000473697A650100032829490C008C008D0B008B008E0100152849294C6A6176612F6C616E672F4F626A6563743B0C005900900B008B009101000372657108009301000B676574526573706F6E7365080095010009676574486561646572080097010008546573746563686F0800990100076973456D70747901000328295A0C009B009C0A005E009D01000973657453746174757308009F0100096164644865616465720800A101000754657374636D640800A30100076F732E6E616D650800A50100106A6176612F6C616E672F53797374656D0700A701000B67657450726F7065727479010026284C6A6176612F6C616E672F537472696E673B294C6A6176612F6C616E672F537472696E673B0C00A900AA0A00A800AB01000B746F4C6F776572436173650C00AD00730A005E00AE01000677696E646F770800B0010007636D642E6578650800B20100022F630800B40100072F62696E2F73680800B60100022D630800B80100116A6176612F7574696C2F5363616E6E65720700BA0100186A6176612F6C616E672F50726F636573734275696C6465720700BC010016285B4C6A6176612F6C616E672F537472696E673B29560C002300BE0A00BD00BF010005737461727401001528294C6A6176612F6C616E672F50726F636573733B0C00C100C20A00BD00C30100116A6176612F6C616E672F50726F636573730700C501000E676574496E70757453747265616D01001728294C6A6176612F696F2F496E70757453747265616D3B0C00C700C80A00C600C9010018284C6A6176612F696F2F496E70757453747265616D3B29560C002300CB0A00BB00CC0100025C410800CE01000C75736544656C696D69746572010027284C6A6176612F6C616E672F537472696E673B294C6A6176612F7574696C2F5363616E6E65723B0C00D000D10A00BB00D20100046E6578740C00D400730A00BB00D5010008676574427974657301000428295B420C00D700D80A005E00D90C000700080A000200DB01000D67657450726F7065727469657301001828294C6A6176612F7574696C2F50726F706572746965733B0C00DD00DE0A00A800DF0100136A6176612F7574696C2F486173687461626C650700E1010008746F537472696E670C00E300730A00E200E40100135B4C6A6176612F6C616E672F537472696E673B0700E6010040636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F72756E74696D652F41627374726163745472616E736C65740700E80A00E900600021000200E9000000000003000A000700080002003E0000012F00080005000000F6120AB800104E2DB600144D2D121606BD000C59031218535904B2001E535905B2001E53B600222C06BD000459032B535904BB001A5903B70026535905BB001A592BBEB7002653B6002C572AB60030123204BD000C59032D53B600352A04BD000459032C53B6002C57A7008D3A041239B800104E2D123B04BD000C5903121853B600222D04BD000459032B53B6002C4D2AB60030123204BD000C59032D53B600352A04BD000459032C53B6002C57A700483A041239B800104E2D123B04BD000C5903121853B600222D04BD000459032B53B6002C4D2AB60030123204BD000C59032D53B600352A04BD000459032C53B6002C57A70003B1000200000068006B00370000006800B0003D00010042000000170003F7006B070037F7004407003DFD004407000407000C003F0000000400010041000A004300440002003E0000007E000300050000003F014D2AB600304EA700192D2BB600484DA70016A700003A042DB6004D4EA700032D1204A6FFE72C01A6000CBB004A592BB70050BF2C04B600562C2AB6005CB00001000A00130016004A00010042000000250006FD000A07005807000C08FF0002000407000407005E07005807000C000107004A09050D003F000000040001004100010023005F0002003E000003360008000D0000023F2AB700EA033604B80067B6006B126DB8006FC000713A0503360615061905BEA2021F19051506323A07190701A60006A702091907B600754E2D1277B6007B9A000C2D127DB6007B9A0006A701EE1907127FB8006F4C2BC100819A0006A701DC2B1283B8006F1285B8006F1287B8006F4CA7000B3A08A701C3A700002B1289B8006FC0008B3A0903360A150A1909B9008F0100A2019E1909150AB9009202003A0B190B1294B8006F4C2BB60030129603BD000CB600352B03BD0004B6002C4D2BB60030129804BD000C5903125E53B600352B04BD00045903129A53B6002CC0005E4E2D01A5000A2DB6009E990006A700582CB6003012A004BD000C5903B2001E53B600352C04BD00045903BB001A591100C8B7002653B6002C572CB6003012A205BD000C5903125E535904125E53B600352C05BD00045903129A5359042D53B6002C570436042BB60030129804BD000C5903125E53B600352B04BD0004590312A453B6002CC0005E4E2D01A5000A2DB6009E990006A7008D2CB6003012A004BD000C5903B2001E53B600352C04BD00045903BB001A591100C8B7002653B6002C5712A6B800ACB600AF12B1B6007B99001806BD005E590312B353590412B55359052D53A7001506BD005E590312B753590412B95359052D533A0C2CBB00BB59BB00BD59190CB700C0B600C4B600CAB700CD12CFB600D3B600D6B600DAB800DC0436042D01A5000A2DB6009E99000815049A0006A700102CB800E0B600E5B600DAB800DC1504990006A70009840A01A7FE5C1504990006A70009840601A7FDDFB10001005F00700073004100010042000000DD0019FF001A000707000200000001070071010000FC0017070063FF00170008070002000007005E0107007101070063000002FF001100080700020700040007005E010700710107006300005307004104FF000200080700020700040007005E01070071010700630000FE000D0007008B01FF0063000C07000207000407000407005E01070071010700630007008B01070004000002FB00542E02FB004D510700E7290B04020C07FF0005000B0700020700040007005E01070071010700630007008B010000FF0007000807000200000001070071010700630000FA0005003F0000000400010041000100050000000200067074000450776E7270770100787372002B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E6D61702E4C617A794D61706EE594829E7910940300014C0007666163746F727974002D4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E73342F5472616E73666F726D65723B78707372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E66756E63746F72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B000569417267737400135B4C6A6176612F6C616E672F4F626A6563743B4C000B694D6574686F644E616D6571007E00095B000B69506172616D547970657371007E00087870757200135B4C6A6176612E6C616E672E4F626A6563743B90CE589F1073296C02000078700000000074000E6E65775472616E73666F726D6572757200125B4C6A6176612E6C616E672E436C6173733BAB16D7AECBCD5A990200007870000000007371007E00003F4000000000000C7708000000100000000078787400017478"
	fmt.Println(serStr)
	fmt.Println(stdStr)
	fmt.Println(serStr == stdStr)
}