$(function () {
    var countries = [
        {
            "area_code": "0093",
            "country_id": 2,
            "name_cn": "阿富汗",
            "name_en": "Afghanistan"
        },
        {
            "area_code": "00355",
            "country_id": 3,
            "name_cn": "阿尔巴尼亚",
            "name_en": "Albania"
        },
        {
            "area_code": "00213",
            "country_id": 4,
            "name_cn": "阿尔及利亚",
            "name_en": "Algeria"
        },
        {
            "area_code": "00376",
            "country_id": 5,
            "name_cn": "安道尔共和国",
            "name_en": "Andorra"
        },
        {
            "area_code": "00244",
            "country_id": 1,
            "name_cn": "安哥拉",
            "name_en": "Angola"
        },
        {
            "area_code": "001264",
            "country_id": 6,
            "name_cn": "安圭拉岛",
            "name_en": "Anguilla"
        },
        {
            "area_code": "001268",
            "country_id": 7,
            "name_cn": "安提瓜和巴布达",
            "name_en": "Antigua and Barbuda"
        },
        {
            "area_code": "0054",
            "country_id": 8,
            "name_cn": "阿根廷",
            "name_en": "Argentina"
        },
        {
            "area_code": "00374",
            "country_id": 9,
            "name_cn": "亚美尼亚",
            "name_en": "Armenia"
        },
        {
            "area_code": "00297",
            "country_id": 194,
            "name_cn": "阿鲁巴",
            "name_en": "Aruba"
        },
        {
            "area_code": "00247",
            "country_id": 10,
            "name_cn": "阿森松",
            "name_en": "Ascension"
        },
        {
            "area_code": "0061",
            "country_id": 11,
            "name_cn": "澳大利亚",
            "name_en": "Australia"
        },
        {
            "area_code": "00672",
            "country_id": 195,
            "name_cn": "澳大利亚海外领地",
            "name_en": "Australian overseas territories"
        },
        {
            "area_code": "0043",
            "country_id": 12,
            "name_cn": "奥地利",
            "name_en": "Austria"
        },
        {
            "area_code": "00994",
            "country_id": 13,
            "name_cn": "阿塞拜疆",
            "name_en": "Azerbaijan"
        },
        {
            "area_code": "001242",
            "country_id": 14,
            "name_cn": "巴哈马",
            "name_en": "Bahamas"
        },
        {
            "area_code": "00973",
            "country_id": 15,
            "name_cn": "巴林",
            "name_en": "Bahrain"
        },
        {
            "area_code": "00880",
            "country_id": 16,
            "name_cn": "孟加拉国",
            "name_en": "Bangladesh"
        },
        {
            "area_code": "001246",
            "country_id": 17,
            "name_cn": "巴巴多斯",
            "name_en": "Barbados"
        },
        {
            "area_code": "00375",
            "country_id": 18,
            "name_cn": "白俄罗斯",
            "name_en": "Belarus"
        },
        {
            "area_code": "0032",
            "country_id": 19,
            "name_cn": "比利时",
            "name_en": "Belgium"
        },
        {
            "area_code": "00501",
            "country_id": 20,
            "name_cn": "伯利兹",
            "name_en": "Belize"
        },
        {
            "area_code": "00229",
            "country_id": 21,
            "name_cn": "贝宁",
            "name_en": "Benin"
        },
        {
            "area_code": "001441",
            "country_id": 22,
            "name_cn": "百慕大群岛",
            "name_en": "Bermuda"
        },
        {
            "area_code": "00975",
            "country_id": 196,
            "name_cn": "不丹",
            "name_en": "Bhutan"
        },
        {
            "area_code": "00591",
            "country_id": 23,
            "name_cn": "玻利维亚",
            "name_en": "Bolivia"
        },
        {
            "area_code": "00387",
            "country_id": 197,
            "name_cn": "波斯尼亚和黑塞哥维那",
            "name_en": "Bosnia and Herzegovina"
        },
        {
            "area_code": "00267",
            "country_id": 24,
            "name_cn": "博茨瓦纳",
            "name_en": "Botswana"
        },
        {
            "area_code": "0055",
            "country_id": 25,
            "name_cn": "巴西",
            "name_en": "Brazil"
        },
        {
            "area_code": "001284",
            "country_id": 229,
            "name_cn": "英属维尔京群岛",
            "name_en": "British Virgin Islands"
        },
        {
            "area_code": "00673",
            "country_id": 26,
            "name_cn": "文莱",
            "name_en": "Brunei"
        },
        {
            "area_code": "00359",
            "country_id": 27,
            "name_cn": "保加利亚",
            "name_en": "Bulgaria"
        },
        {
            "area_code": "00226",
            "country_id": 28,
            "name_cn": "布基纳法索",
            "name_en": "Burkina Faso"
        },
        {
            "area_code": "00257",
            "country_id": 30,
            "name_cn": "布隆迪",
            "name_en": "Burundi"
        },
        {
            "area_code": "00855",
            "country_id": 85,
            "name_cn": "柬埔寨",
            "name_en": "Cambodia"
        },
        {
            "area_code": "00237",
            "country_id": 31,
            "name_cn": "喀麦隆",
            "name_en": "Cameroon"
        },
        {
            "area_code": "001",
            "country_id": 32,
            "name_cn": "加拿大",
            "name_en": "Canada"
        },
        {
            "area_code": "00238",
            "country_id": 198,
            "name_cn": "佛得角",
            "name_en": "Cape Verde"
        },
        {
            "area_code": "001345",
            "country_id": 33,
            "name_cn": "开曼群岛",
            "name_en": "Cayman Islands."
        },
        {
            "area_code": "00236",
            "country_id": 34,
            "name_cn": "中非共和国",
            "name_en": "Central African Republic"
        },
        {
            "area_code": "00235",
            "country_id": 35,
            "name_cn": "乍得",
            "name_en": "Chad"
        },
        {
            "area_code": "0056",
            "country_id": 36,
            "name_cn": "智利",
            "name_en": "Chile"
        },
        {
            "area_code": "0086",
            "country_id": 37,
            "name_cn": "中国",
            "name_en": "China"
        },
        {
            "area_code": "0057",
            "country_id": 38,
            "name_cn": "哥伦比亚",
            "name_en": "Colombia"
        },
        {
            "area_code": "00269",
            "country_id": 199,
            "name_cn": "科摩罗群岛",
            "name_en": "Comoros Islands"
        },
        {
            "area_code": "00243",
            "country_id": 39,
            "name_cn": "刚果",
            "name_en": "Congo"
        },
        {
            "area_code": "00682",
            "country_id": 40,
            "name_cn": "库克群岛",
            "name_en": "Cook Islands."
        },
        {
            "area_code": "00506",
            "country_id": 41,
            "name_cn": "哥斯达黎加",
            "name_en": "Costa Rica"
        },
        {
            "area_code": "00385",
            "country_id": 200,
            "name_cn": "克罗地亚",
            "name_en": "Croatia"
        },
        {
            "area_code": "0053",
            "country_id": 42,
            "name_cn": "古巴",
            "name_en": "Cuba"
        },
        {
            "area_code": "00357",
            "country_id": 43,
            "name_cn": "塞浦路斯",
            "name_en": "Cyprus"
        },
        {
            "area_code": "00420",
            "country_id": 44,
            "name_cn": "捷克",
            "name_en": "Czech Republic"
        },
        {
            "area_code": "0045",
            "country_id": 45,
            "name_cn": "丹麦",
            "name_en": "Denmark"
        },
        {
            "area_code": "00246",
            "country_id": 201,
            "name_cn": "迭戈加西亚群岛",
            "name_en": "Diego Garcia"
        },
        {
            "area_code": "00253",
            "country_id": 46,
            "name_cn": "吉布提",
            "name_en": "Djibouti"
        },
        {
            "area_code": "001809",
            "country_id": 47,
            "name_cn": "多米尼加共和国",
            "name_en": "Dominican Republic"
        },
        {
            "area_code": "00670",
            "country_id": 202,
            "name_cn": "东帝汶",
            "name_en": "East Timor"
        },
        {
            "area_code": "00593",
            "country_id": 48,
            "name_cn": "厄瓜多尔",
            "name_en": "Ecuador"
        },
        {
            "area_code": "0020",
            "country_id": 49,
            "name_cn": "埃及",
            "name_en": "Egypt"
        },
        {
            "area_code": "00503",
            "country_id": 50,
            "name_cn": "萨尔瓦多",
            "name_en": "El Salvador"
        },
        {
            "area_code": "009714",
            "country_id": 230,
            "name_cn": "迪拜酋长国",
            "name_en": "Emirate of Dubai"
        },
        {
            "area_code": "00240",
            "country_id": 203,
            "name_cn": "赤道几内亚",
            "name_en": "Equatorial Guinea"
        },
        {
            "area_code": "00291",
            "country_id": 204,
            "name_cn": "厄立特里亚",
            "name_en": "Eritrea"
        },
        {
            "area_code": "00372",
            "country_id": 51,
            "name_cn": "爱沙尼亚",
            "name_en": "Estonia"
        },
        {
            "area_code": "00251",
            "country_id": 52,
            "name_cn": "埃塞俄比亚",
            "name_en": "Ethiopia"
        },
        {
            "area_code": "00500",
            "country_id": 205,
            "name_cn": "福克兰群岛",
            "name_en": "Falkland Islands"
        },
        {
            "area_code": "00298",
            "country_id": 206,
            "name_cn": "法罗群岛",
            "name_en": "Faroe Islands"
        },
        {
            "area_code": "00679",
            "country_id": 53,
            "name_cn": "斐济",
            "name_en": "Fiji"
        },
        {
            "area_code": "00358",
            "country_id": 54,
            "name_cn": "芬兰",
            "name_en": "Finland"
        },
        {
            "area_code": "0033",
            "country_id": 55,
            "name_cn": "法国",
            "name_en": "France"
        },
        {
            "area_code": "00594",
            "country_id": 56,
            "name_cn": "法属圭亚那",
            "name_en": "French Guiana"
        },
        {
            "area_code": "00689",
            "country_id": 136,
            "name_cn": "法属玻利尼西亚",
            "name_en": "French Polynesia"
        },
        {
            "area_code": "00241",
            "country_id": 57,
            "name_cn": "加蓬",
            "name_en": "Gabon"
        },
        {
            "area_code": "00220",
            "country_id": 58,
            "name_cn": "冈比亚",
            "name_en": "Gambia"
        },
        {
            "area_code": "00995",
            "country_id": 59,
            "name_cn": "格鲁吉亚",
            "name_en": "Georgia"
        },
        {
            "area_code": "0049",
            "country_id": 60,
            "name_cn": "德国",
            "name_en": "Germany"
        },
        {
            "area_code": "00233",
            "country_id": 61,
            "name_cn": "加纳",
            "name_en": "Ghana"
        },
        {
            "area_code": "00350",
            "country_id": 62,
            "name_cn": "直布罗陀",
            "name_en": "Gibraltar"
        },
        {
            "area_code": "0030",
            "country_id": 63,
            "name_cn": "希腊",
            "name_en": "Greece"
        },
        {
            "area_code": "00299",
            "country_id": 207,
            "name_cn": "格陵兰岛",
            "name_en": "Greenland"
        },
        {
            "area_code": "001473",
            "country_id": 64,
            "name_cn": "格林纳达",
            "name_en": "Grenada"
        },
        {
            "area_code": "00590",
            "country_id": 208,
            "name_cn": "瓜德罗普",
            "name_en": "Guadeloupe"
        },
        {
            "area_code": "001671",
            "country_id": 65,
            "name_cn": "关岛",
            "name_en": "Guam"
        },
        {
            "area_code": "00502",
            "country_id": 66,
            "name_cn": "危地马拉",
            "name_en": "Guatemala"
        },
        {
            "area_code": "00224",
            "country_id": 67,
            "name_cn": "几内亚",
            "name_en": "Guinea"
        },
        {
            "area_code": "00245",
            "country_id": 209,
            "name_cn": "几内亚比绍",
            "name_en": "Guinea-Bissau"
        },
        {
            "area_code": "00592",
            "country_id": 68,
            "name_cn": "圭亚那",
            "name_en": "Guyana"
        },
        {
            "area_code": "00509",
            "country_id": 69,
            "name_cn": "海地",
            "name_en": "Haiti"
        },
        {
            "area_code": "00504",
            "country_id": 70,
            "name_cn": "洪都拉斯",
            "name_en": "Honduras"
        },
        {
            "area_code": "00852",
            "country_id": 71,
            "name_cn": "中国香港",
            "name_en": "Hong Kong (China)"
        },
        {
            "area_code": "0036",
            "country_id": 72,
            "name_cn": "匈牙利",
            "name_en": "Hungary"
        },
        {
            "area_code": "00354",
            "country_id": 73,
            "name_cn": "冰岛",
            "name_en": "Iceland"
        },
        {
            "area_code": "0091",
            "country_id": 74,
            "name_cn": "印度",
            "name_en": "India"
        },
        {
            "area_code": "0062",
            "country_id": 75,
            "name_cn": "印度尼西亚",
            "name_en": "Indonesia"
        },
        {
            "area_code": "0098",
            "country_id": 76,
            "name_cn": "伊朗",
            "name_en": "Iran"
        },
        {
            "area_code": "00964",
            "country_id": 77,
            "name_cn": "伊拉克",
            "name_en": "Iraq"
        },
        {
            "area_code": "00353",
            "country_id": 78,
            "name_cn": "爱尔兰",
            "name_en": "Ireland"
        },
        {
            "area_code": "00972",
            "country_id": 79,
            "name_cn": "以色列",
            "name_en": "Israel"
        },
        {
            "area_code": "0039",
            "country_id": 80,
            "name_cn": "意大利",
            "name_en": "Italy"
        },
        {
            "area_code": "00225",
            "country_id": 81,
            "name_cn": "科特迪瓦",
            "name_en": "Ivory Coast"
        },
        {
            "area_code": "001876",
            "country_id": 82,
            "name_cn": "牙买加",
            "name_en": "Jamaica"
        },
        {
            "area_code": "0081",
            "country_id": 83,
            "name_cn": "日本",
            "name_en": "Japan"
        },
        {
            "area_code": "00962",
            "country_id": 84,
            "name_cn": "约旦",
            "name_en": "Jordan"
        },
        {
            "area_code": "007",
            "country_id": 86,
            "name_cn": "哈萨克斯坦",
            "name_en": "Kazakhstan"
        },
        {
            "area_code": "00254",
            "country_id": 87,
            "name_cn": "肯尼亚",
            "name_en": "Kenya"
        },
        {
            "area_code": "00686",
            "country_id": 210,
            "name_cn": "基里巴斯",
            "name_en": "Kiribati"
        },
        {
            "area_code": "0082",
            "country_id": 88,
            "name_cn": "韩国",
            "name_en": "Korea"
        },
        {
            "area_code": "00965",
            "country_id": 89,
            "name_cn": "科威特",
            "name_en": "Kuwait"
        },
        {
            "area_code": "00996",
            "country_id": 90,
            "name_cn": "吉尔吉斯坦",
            "name_en": "Kyrgyzstan"
        },
        {
            "area_code": "00856",
            "country_id": 91,
            "name_cn": "老挝",
            "name_en": "Laos"
        },
        {
            "area_code": "00371",
            "country_id": 92,
            "name_cn": "拉脱维亚",
            "name_en": "Latvia"
        },
        {
            "area_code": "00961",
            "country_id": 93,
            "name_cn": "黎巴嫩",
            "name_en": "Lebanon"
        },
        {
            "area_code": "00266",
            "country_id": 94,
            "name_cn": "莱索托",
            "name_en": "Lesotho"
        },
        {
            "area_code": "00231",
            "country_id": 95,
            "name_cn": "利比里亚",
            "name_en": "Liberia"
        },
        {
            "area_code": "00218",
            "country_id": 96,
            "name_cn": "利比亚",
            "name_en": "Libya"
        },
        {
            "area_code": "00423",
            "country_id": 97,
            "name_cn": "列支敦士登",
            "name_en": "Liechtenstein"
        },
        {
            "area_code": "00370",
            "country_id": 98,
            "name_cn": "立陶宛",
            "name_en": "Lithuania"
        },
        {
            "area_code": "00352",
            "country_id": 99,
            "name_cn": "卢森堡",
            "name_en": "Luxembourg"
        },
        {
            "area_code": "00389",
            "country_id": 211,
            "name_cn": "马其顿",
            "name_en": "Macedonia"
        },
        {
            "area_code": "00261",
            "country_id": 101,
            "name_cn": "马达加斯加",
            "name_en": "Madagascar"
        },
        {
            "area_code": "00265",
            "country_id": 102,
            "name_cn": "马拉维",
            "name_en": "Malawi"
        },
        {
            "area_code": "0060",
            "country_id": 103,
            "name_cn": "马来西亚",
            "name_en": "Malaysia"
        },
        {
            "area_code": "00960",
            "country_id": 104,
            "name_cn": "马尔代夫",
            "name_en": "Maldives"
        },
        {
            "area_code": "00223",
            "country_id": 105,
            "name_cn": "马里",
            "name_en": "Mali"
        },
        {
            "area_code": "00356",
            "country_id": 106,
            "name_cn": "马耳他",
            "name_en": "Malta"
        },
        {
            "area_code": "00223",
            "country_id": 107,
            "name_cn": "马里亚那群岛",
            "name_en": "Mariana Islands"
        },
        {
            "area_code": "00692",
            "country_id": 212,
            "name_cn": "马绍尔群岛",
            "name_en": "Marshall Islands"
        },
        {
            "area_code": "00596",
            "country_id": 108,
            "name_cn": "马提尼克",
            "name_en": "Martinique"
        },
        {
            "area_code": "00222",
            "country_id": 213,
            "name_cn": "毛里塔尼亚",
            "name_en": "Mauritania"
        },
        {
            "area_code": "00230",
            "country_id": 109,
            "name_cn": "毛里求斯",
            "name_en": "Mauritius"
        },
        {
            "area_code": "0052",
            "country_id": 110,
            "name_cn": "墨西哥",
            "name_en": "Mexico"
        },
        {
            "area_code": "00691",
            "country_id": 214,
            "name_cn": "密克罗尼西亚",
            "name_en": "Micronesia"
        },
        {
            "area_code": "00373",
            "country_id": 111,
            "name_cn": "摩尔多瓦",
            "name_en": "Moldova"
        },
        {
            "area_code": "00377",
            "country_id": 112,
            "name_cn": "摩纳哥",
            "name_en": "Monaco"
        },
        {
            "area_code": "00976",
            "country_id": 113,
            "name_cn": "蒙古",
            "name_en": "Mongolia"
        },
        {
            "area_code": "00382",
            "country_id": 215,
            "name_cn": "黑山",
            "name_en": "Montenegro"
        },
        {
            "area_code": "001664",
            "country_id": 114,
            "name_cn": "蒙特塞拉特岛",
            "name_en": "Montserrat"
        },
        {
            "area_code": "00212",
            "country_id": 115,
            "name_cn": "摩洛哥",
            "name_en": "Morocco"
        },
        {
            "area_code": "00258",
            "country_id": 116,
            "name_cn": "莫桑比克",
            "name_en": "Mozambique"
        },
        {
            "area_code": "0095",
            "country_id": 29,
            "name_cn": "缅甸",
            "name_en": "Myanmar"
        },
        {
            "area_code": "00264",
            "country_id": 117,
            "name_cn": "纳米比亚",
            "name_en": "Namibia"
        },
        {
            "area_code": "00674",
            "country_id": 118,
            "name_cn": "瑙鲁",
            "name_en": "Nauru"
        },
        {
            "area_code": "00977",
            "country_id": 119,
            "name_cn": "尼泊尔",
            "name_en": "Nepal"
        },
        {
            "area_code": "00599",
            "country_id": 120,
            "name_cn": "荷属安的列斯",
            "name_en": "Netheriands Antilles"
        },
        {
            "area_code": "0031",
            "country_id": 121,
            "name_cn": "荷兰",
            "name_en": "Netherlands"
        },
        {
            "area_code": "00687",
            "country_id": 216,
            "name_cn": "新喀里多尼亚",
            "name_en": "New Caledonia"
        },
        {
            "area_code": "0064",
            "country_id": 122,
            "name_cn": "新西兰",
            "name_en": "New Zealand"
        },
        {
            "area_code": "00505",
            "country_id": 123,
            "name_cn": "尼加拉瓜",
            "name_en": "Nicaragua"
        },
        {
            "area_code": "00227",
            "country_id": 124,
            "name_cn": "尼日尔",
            "name_en": "Niger"
        },
        {
            "area_code": "00234",
            "country_id": 125,
            "name_cn": "尼日利亚",
            "name_en": "Nigeria"
        },
        {
            "area_code": "00683",
            "country_id": 217,
            "name_cn": "纽埃岛",
            "name_en": "Niue"
        },
        {
            "area_code": "00850",
            "country_id": 126,
            "name_cn": "朝鲜",
            "name_en": "North Korea"
        },
        {
            "area_code": "0047",
            "country_id": 127,
            "name_cn": "挪威",
            "name_en": "Norway"
        },
        {
            "area_code": "00968",
            "country_id": 128,
            "name_cn": "阿曼",
            "name_en": "Oman"
        },
        {
            "area_code": "0092",
            "country_id": 129,
            "name_cn": "巴基斯坦",
            "name_en": "Pakistan"
        },
        {
            "area_code": "00680",
            "country_id": 218,
            "name_cn": "帕劳",
            "name_en": "Palau"
        },
        {
            "area_code": "00970",
            "country_id": 219,
            "name_cn": "巴勒斯坦",
            "name_en": "Palestine"
        },
        {
            "area_code": "00507",
            "country_id": 130,
            "name_cn": "巴拿马",
            "name_en": "Panama"
        },
        {
            "area_code": "00675",
            "country_id": 131,
            "name_cn": "巴布亚新几内亚",
            "name_en": "Papua New Guinea"
        },
        {
            "area_code": "00595",
            "country_id": 132,
            "name_cn": "巴拉圭",
            "name_en": "Paraguay"
        },
        {
            "area_code": "0051",
            "country_id": 133,
            "name_cn": "秘鲁",
            "name_en": "Peru"
        },
        {
            "area_code": "0063",
            "country_id": 134,
            "name_cn": "菲律宾",
            "name_en": "Philippines"
        },
        {
            "area_code": "0048",
            "country_id": 135,
            "name_cn": "波兰",
            "name_en": "Poland"
        },
        {
            "area_code": "00351",
            "country_id": 137,
            "name_cn": "葡萄牙",
            "name_en": "Portugal"
        },
        {
            "area_code": "001",
            "country_id": 138,
            "name_cn": "波多黎各",
            "name_en": "Puerto Rico"
        },
        {
            "area_code": "00974",
            "country_id": 139,
            "name_cn": "卡塔尔",
            "name_en": "Qatar"
        },
        {
            "area_code": "00262",
            "country_id": 140,
            "name_cn": "留尼旺",
            "name_en": "Reunion"
        },
        {
            "area_code": "0040",
            "country_id": 141,
            "name_cn": "罗马尼亚",
            "name_en": "Romania"
        },
        {
            "area_code": "007",
            "country_id": 142,
            "name_cn": "俄罗斯",
            "name_en": "Russia"
        },
        {
            "area_code": "00250",
            "country_id": 220,
            "name_cn": "卢旺达",
            "name_en": "Rwanda"
        },
        {
            "area_code": "001758",
            "country_id": 143,
            "name_cn": "圣卢西亚",
            "name_en": "Saint Lucia"
        },
        {
            "area_code": "00508",
            "country_id": 222,
            "name_cn": "圣皮埃尔和密克隆群岛",
            "name_en": "Saint Pierre and Miquelon"
        },
        {
            "area_code": "001784",
            "country_id": 144,
            "name_cn": "圣文森特岛",
            "name_en": "Saint Vincent"
        },
        {
            "area_code": "00684",
            "country_id": 145,
            "name_cn": "东萨摩亚(美)",
            "name_en": "Samoa Eastern"
        },
        {
            "area_code": "00685",
            "country_id": 146,
            "name_cn": "西萨摩亚",
            "name_en": "Samoa Western"
        },
        {
            "area_code": "00378",
            "country_id": 147,
            "name_cn": "圣马力诺",
            "name_en": "San Marino"
        },
        {
            "area_code": "00239",
            "country_id": 148,
            "name_cn": "圣多美和普林西比",
            "name_en": "Sao Tome and Principe"
        },
        {
            "area_code": "00966",
            "country_id": 149,
            "name_cn": "沙特阿拉伯",
            "name_en": "Saudi Arabia"
        },
        {
            "area_code": "00221",
            "country_id": 150,
            "name_cn": "塞内加尔",
            "name_en": "Senegal"
        },
        {
            "area_code": "00381",
            "country_id": 223,
            "name_cn": "塞尔维亚",
            "name_en": "Serbia"
        },
        {
            "area_code": "00248",
            "country_id": 151,
            "name_cn": "塞舌尔",
            "name_en": "Seychelles"
        },
        {
            "area_code": "00232",
            "country_id": 152,
            "name_cn": "塞拉利昂",
            "name_en": "Sierra Leone"
        },
        {
            "area_code": "0065",
            "country_id": 153,
            "name_cn": "新加坡",
            "name_en": "Singapore"
        },
        {
            "area_code": "00421",
            "country_id": 154,
            "name_cn": "斯洛伐克",
            "name_en": "Slovakia"
        },
        {
            "area_code": "00386",
            "country_id": 155,
            "name_cn": "斯洛文尼亚",
            "name_en": "Slovenia"
        },
        {
            "area_code": "00677",
            "country_id": 156,
            "name_cn": "所罗门群岛",
            "name_en": "Solomon Islands"
        },
        {
            "area_code": "00252",
            "country_id": 157,
            "name_cn": "索马里",
            "name_en": "Somalia"
        },
        {
            "area_code": "0027",
            "country_id": 158,
            "name_cn": "南非",
            "name_en": "South Africa"
        },
        {
            "area_code": "0027",
            "country_id": 190,
            "name_cn": "南非",
            "name_en": "South Africa"
        },
        {
            "area_code": "0034",
            "country_id": 159,
            "name_cn": "西班牙",
            "name_en": "Spain"
        },
        {
            "area_code": "0094",
            "country_id": 160,
            "name_cn": "斯里兰卡",
            "name_en": "Sri Lanka"
        },
        {
            "area_code": "00290",
            "country_id": 221,
            "name_cn": "圣赫勒拿岛",
            "name_en": "St.Helena"
        },
        {
            "area_code": "001758",
            "country_id": 161,
            "name_cn": "圣卢西亚",
            "name_en": "St.Lucia"
        },
        {
            "area_code": "001784",
            "country_id": 162,
            "name_cn": "圣文森特",
            "name_en": "St.Vincent"
        },
        {
            "area_code": "00249",
            "country_id": 163,
            "name_cn": "苏丹",
            "name_en": "Sudan"
        },
        {
            "area_code": "00597",
            "country_id": 164,
            "name_cn": "苏里南",
            "name_en": "Suriname"
        },
        {
            "area_code": "00268",
            "country_id": 165,
            "name_cn": "斯威士兰",
            "name_en": "Swaziland"
        },
        {
            "area_code": "0046",
            "country_id": 166,
            "name_cn": "瑞典",
            "name_en": "Sweden"
        },
        {
            "area_code": "0041",
            "country_id": 167,
            "name_cn": "瑞士",
            "name_en": "Switzerland"
        },
        {
            "area_code": "00963",
            "country_id": 168,
            "name_cn": "叙利亚",
            "name_en": "Syria"
        },
        {
            "area_code": "00886",
            "country_id": 169,
            "name_cn": "中国台湾",
            "name_en": "Taiwan (China)"
        },
        {
            "area_code": "00992",
            "country_id": 170,
            "name_cn": "塔吉克斯坦",
            "name_en": "Tajikistan"
        },
        {
            "area_code": "00255",
            "country_id": 171,
            "name_cn": "坦桑尼亚",
            "name_en": "Tanzania"
        },
        {
            "area_code": "0066",
            "country_id": 172,
            "name_cn": "泰国",
            "name_en": "Thailand"
        },
        {
            "area_code": "00228",
            "country_id": 173,
            "name_cn": "多哥",
            "name_en": "Togo"
        },
        {
            "area_code": "00690",
            "country_id": 224,
            "name_cn": "托克劳群岛",
            "name_en": "Tokelau"
        },
        {
            "area_code": "00676",
            "country_id": 174,
            "name_cn": "汤加",
            "name_en": "Tonga"
        },
        {
            "area_code": "001868",
            "country_id": 175,
            "name_cn": "特立尼达和多巴哥",
            "name_en": "Trinidad and Tobago"
        },
        {
            "area_code": "00216",
            "country_id": 176,
            "name_cn": "突尼斯",
            "name_en": "Tunisia"
        },
        {
            "area_code": "0090",
            "country_id": 177,
            "name_cn": "土耳其",
            "name_en": "Turkey"
        },
        {
            "area_code": "00993",
            "country_id": 178,
            "name_cn": "土库曼斯坦",
            "name_en": "Turkmenistan"
        },
        {
            "area_code": "00688",
            "country_id": 225,
            "name_cn": "图瓦卢",
            "name_en": "Tuvalu"
        },
        {
            "area_code": "00256",
            "country_id": 179,
            "name_cn": "乌干达",
            "name_en": "Uganda"
        },
        {
            "area_code": "00380",
            "country_id": 180,
            "name_cn": "乌克兰",
            "name_en": "Ukraine"
        },
        {
            "area_code": "00971",
            "country_id": 181,
            "name_cn": "阿拉伯联合酋长国",
            "name_en": "United Arab Emirates"
        },
        {
            "area_code": "0044",
            "country_id": 182,
            "name_cn": "英国",
            "name_en": "United Kingdom"
        },
        {
            "area_code": "001",
            "country_id": 183,
            "name_cn": "美国",
            "name_en": "United States of America"
        },
        {
            "area_code": "00598",
            "country_id": 184,
            "name_cn": "乌拉圭",
            "name_en": "Uruguay"
        },
        {
            "area_code": "00998",
            "country_id": 185,
            "name_cn": "乌兹别克斯坦",
            "name_en": "Uzbekistan"
        },
        {
            "area_code": "00678",
            "country_id": 226,
            "name_cn": "瓦努阿图",
            "name_en": "Vanuatu"
        },
        {
            "area_code": "00379",
            "country_id": 227,
            "name_cn": "梵蒂冈城",
            "name_en": "Vatican City"
        },
        {
            "area_code": "0058",
            "country_id": 186,
            "name_cn": "委内瑞拉",
            "name_en": "Venezuela"
        },
        {
            "area_code": "0084",
            "country_id": 187,
            "name_cn": "越南",
            "name_en": "Vietnam"
        },
        {
            "area_code": "00681",
            "country_id": 228,
            "name_cn": "瓦利斯和富图纳",
            "name_en": "Wallis and Futuna"
        },
        {
            "area_code": "00967",
            "country_id": 188,
            "name_cn": "也门",
            "name_en": "Yemen"
        },
        {
            "area_code": "00338",
            "country_id": 189,
            "name_cn": "南斯拉夫",
            "name_en": "Yugoslavia"
        },
        {
            "area_code": "00243",
            "country_id": 192,
            "name_cn": "扎伊尔",
            "name_en": "Zaire"
        },
        {
            "area_code": "00260",
            "country_id": 193,
            "name_cn": "赞比亚",
            "name_en": "Zambia"
        },
        {
            "area_code": "00263",
            "country_id": 191,
            "name_cn": "津巴布韦",
            "name_en": "Zimbabwe"
        }
    ];

    var countryInput = $(document).find('.countrypicker');
    var countryList = "";


    //set defaults
    for (i = 0; i < countryInput.length; i++) {

        //check if flag
        flag = countryInput.eq(i).data('flag');

        if (flag) {
            countryList = "";

            //for each build list with flag
            $.each(countries, function (index, country) {
                var flagIcon = "css/flags/" + country.code + ".png";
                countryList += "<option data-country-code='" + country.area_code + "' data-tokens='" + country.area_code + " " + country.area_code + "' style='padding-left:25px; background-position: 4px 7px; background-image:url(" + flagIcon + ");background-repeat:no-repeat;' value='" + country.area_code + "'>" + country.en_name + "</option>";
            });

            //add flags to button
            countryInput.eq(i).on('loaded.bs.select', function (e) {
                var button = $(this).closest('.btn-group').children('.btn');
                button.hide();
                var def = $(this).find(':selected').data('country-code');
                var flagIcon = "css/flags/" + def + ".png";
                button.css("background-size", '20px');
                button.css("background-position", '10px 9px');
                button.css("padding-left", '40px');
                button.css("background-repeat", 'no-repeat');
                button.css("background-image", "url('" + flagIcon + "'");
                button.show();
            });

            //change flag on select change
            countryInput.eq(i).on('change', function () {
                button = $(this).closest('.btn-group').children('.btn');
                def = $(this).find(':selected').data('country-code');
                flagIcon = "css/flags/" + def + ".png";
                button.css("background-size", '20px');
                button.css("background-position", '10px 9px');
                button.css("padding-left", '40px');
                button.css("background-repeat", 'no-repeat');
                button.css("background-image", "url('" + flagIcon + "'");

            });
        } else {
            countryList = "";

            //for each build list without flag
            $.each(countries, function (index, country) {
                countryList += "<option data-country-code='" + country.area_code + "' data-tokens='" + country.area_code + " " + country.area_code + "' value='" + country.area_code + "'>" + country.en_name + "</option>";
            });


        }

        //append country list
        countryInput.eq(i).html(countryList);


        //check if default
        def = countryInput.eq(i).data('default');
        //if there's a default, set it
        if (def) {
            countryInput.eq(i).val(def);
        }


    }


});