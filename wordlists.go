package moderation

// Generated by generator/filter_dictionary.go; DO NOT EDIT

var profanities = []string{"anal", "anus", "arse", "ass", "balls", "ballsack", "bastard", "biatch", "bitch", "blowjob", "bollock", "bollok", "boner", "boob", "btch", "bugger", "butt", "clitoris", "cock", "coon", "crap", "cum", "cunt", "dick", "dildo", "dyke", "fag", "feck", "felching", "fellate", "fellatio", "flange", "fuck", "fudgepacker", "horny", "incest", "jerk", "jizz", "labia", "muff", "naked", "nigga", "nigger", "nude", "penis", "piss", "poop", "porn", "prick", "prostitut", "pube", "pussy", "queer", "rimjob", "scrotum", "semen", "sex", "shit", "slut", "spunk", "suckmy", "tit", "tohell", "tosser", "turd", "twat", "vagina", "wank", "whore"}

var falsePositives = []string{"abricock", "accum", "acock", "acolapissa", "acrosarcum", "acumen", "agapornis", "akshita", "alcumy", "alexipharmacum", "alkanal", "altitonant", "altitude", "altitudinal", "altitudinarian", "altitudinous", "amatito", "ambosexous", "ambosexual", "ammoniacum", "anacleticum", "analabos", "analagous", "analav", "analcime", "analcimic", "analcimite", "analcite", "analecta", "analectic", "analects", "analemma", "analepses", "analepsis", "analepsy", "analeptic", "analgen", "analgesia", "analgesic", "analgesidae", "analgesis", "analgetic", "analgia", "analgic", "analgize", "analities", "anality", "analkalinity", "anallagmatic", "anallagmatis", "anallantoic", "anallantoidea", "anallergic", "anally", "analog", "analphabet", "analysability", "analysable", "analysand", "analysation", "analyse", "analysing", "analysis", "analyst", "analyt", "analyzability", "analyzable", "analyzation", "analyze", "analyzing", "angraecum", "antesignanus", "anthericum", "anticum", "antisex", "antitabetic", "antitabloid", "antitax", "antithrombic", "antithrombin", "antithyroid", "antitobacco", "antitonic", "antitorpedo", "antitoxic", "antitoxin", "antitrade", "antitradition", "antitragal", "antitragi", "antitragus", "antitrinitarian", "antitrismus", "antitrochanter", "antitropal", "antitrope", "antitropic", "antitropous", "antitropy", "antitrust", "antitrypsin", "antitryptic", "antitubercular", "antituberculin", "antituberculosis", "antituberculotic", "antituberculous", "antitumor", "antiturnpikeism", "antitwilight", "antitypal", "antitype", "antityphoid", "antitypic", "antitypous", "antitypy", "antityrosinase", "anusim", "anusvara", "apoop", "aprosexia", "aptitude", "aptitudinal", "arctitude", "ardass", "aromatitae", "arsedine", "arsefoot", "arsehole", "arsenal", "arsenate", "arsenation", "arseneted", "arsenetted", "arsenfast", "arsenferratose", "arsenhemol", "arseniasis", "arseniate", "arsenic", "arsenide", "arseniferous", "arsenillo", "arseniopleite", "arseniosiderite", "arsenious", "arsenism", "arsenite", "arsenium", "arseniuret", "arsenization", "arseno", "arsenyl", "arsesmart", "artisanal", "asexual", "assacu", "assafetida", "assafoetida", "assagai", "assahy", "assai", "assalto", "assam", "assapan", "assarion", "assart", "assary", "assate", "assation", "assaugement", "assault", "assausive", "assaut", "assay", "assbaa", "asse", "asshead", "asshole", "assi", "asslike", "assman", "assn", "assobre", "assoc", "assoil", "assoin", "assoluto", "assonance", "assonant", "assonate", "assonia", "assoria", "assort", "assot", "asssembler", "asst", "assuade", "assuagable", "assuage", "assuaging", "assuasive", "assubjugate", "assuefaction", "assuetude", "assumable", "assumably", "assume", "assuming", "assummon", "assumpsit", "assumpt", "assurable", "assurance", "assurant", "assurate", "assurd", "assure", "assurge", "assuring", "assuror", "asswage", "asswaging", "assyria", "assyriological", "assyriologist", "assyriologue", "assyriology", "assyroid", "assyth", "asylabia", "asyllabia", "attitude", "attitudinal", "attitudinarian", "attitudinise", "attitudinising", "attitudinize", "attitudinizing", "attitudist", "aunjetitz", "aurichalcum", "autosexing", "azelfafage", "babcock", "bacchanal", "badass", "bagass", "balanus", "ballsier", "ballsiest", "ballstock", "ballsy", "banal", "bananaland", "banus", "barse", "baseballs", "basemen", "basilicock", "basketballs", "bass", "bawcock", "beanballs", "beatitude", "begass", "benedick", "bennettitaceae", "bennettitaceous", "bennettitales", "bereshith", "bewhore", "bibcock", "bidcock", "bilcock", "billycock", "biplanal", "bisexed", "bisext", "bisexual", "bisexuous", "bissext", "bitched", "bitcheries", "bitchery", "bitchier", "bitchiest", "bitchily", "bitchiness", "bitching", "bitchy", "blackballs", "blackbutt", "blackcock", "blowballs", "blowcock", "bluetit", "bonassus", "boobery", "boobialla", "boobily", "boobish", "booboisie", "booboo", "boobyalla", "boobyish", "boobyism", "borassus", "bouffage", "bountith", "brainfag", "brass", "brushite", "buckass", "bufagin", "buggered", "buggeries", "buggering", "buggery", "bushtit", "butanal", "buttal", "butte", "buttgenbachite", "butthe", "butties", "butting", "buttinski", "buttinsky", "buttle", "buttling", "buttock", "button", "buttress", "buttstock", "buttstrap", "buttwoman", "buttwomen", "buttwood", "butty", "cacoon", "caecum", "canal", "canarsee", "cannonballs", "canvass", "capsicum", "carassow", "carbonero", "carcoon", "carse", "cass", "castellanus", "catacumba", "catharses", "cathedraticum", "cecum", "celestitude", "ceratitoid", "certitude", "cespititous", "chalcidicum", "chashitsu", "chass", "chastity", "chauffage", "chercock", "chickenshit", "chiefage", "chlorococcum", "chrysophanus", "chuprassy", "cimcumvention", "circum", "cirmcumferential", "clarseach", "clarsech", "clarseth", "clericum", "clinchpoop", "clitorism", "coarse", "cockabondy", "cockade", "cockadoodledoo", "cockaigne", "cockal", "cockamamie", "cockamamy", "cockamaroo", "cockandy", "cockapoo", "cockard", "cockarouse", "cockateel", "cockatiel", "cockatoo", "cockatrice", "cockawee", "cockbell", "cockbill", "cockbird", "cockboat", "cockbrain", "cockburn", "cockchafer", "cockcrow", "cocked", "cocker", "cocket", "cockeye", "cockfight", "cockhead", "cockhorse", "cockie", "cockily", "cockiness", "cocking", "cockish", "cockle", "cocklight", "cocklike", "cockling", "cockloche", "cockloft", "cockly", "cockmatch", "cockmate", "cockneian", "cockneity", "cockney", "cockpaddle", "cockpit", "cockroach", "cockscomb", "cocksfoot", "cockshead", "cockshies", "cockshoot", "cockshot", "cockshut", "cockshy", "cocksparrow", "cockspur", "cockstone", "cocksure", "cockswain", "cocksy", "cocktail", "cockthrowing", "cockup", "cockweed", "cocky", "cocoon", "coecum", "colchicum", "coldcock", "coletit", "committitur", "concumbency", "constituencies", "constituency", "constituent", "constitute", "constituting", "constitution", "constitutive", "constitutor", "consuetitude", "contrafagotto", "cooncan", "cooner", "coonhound", "coonier", "cooniest", "coonily", "cooniness", "coonjine", "coonroot", "coonskin", "coontah", "coontail", "coontie", "coony", "coriolanus", "cornballs", "counterflange", "counterprick", "coverslut", "crapaud", "crape", "craping", "crapon", "crapped", "crappin", "crapple", "crappo", "crapshooter", "crapshooting", "crapula", "crapulence", "crapulency", "crapulent", "crapulous", "crapwa", "crass", "crevass", "cucumiform", "cucumis", "cuirass", "cumacea", "cumaceous", "cumaean", "cumal", "cumanagoto", "cumaphyte", "cumaphytic", "cumaphytism", "cumar", "cumay", "cumbent", "cumber", "cumbha", "cumble", "cumbly", "cumbraite", "cumbrance", "cumbre", "cumbrian", "cumbrous", "cumbu", "cumene", "cumengite", "cumenyl", "cumflutter", "cumhal", "cumic", "cumidin", "cumin", "cumly", "cummer", "cummin", "cummock", "cumol", "cump", "cumquat", "cumsha", "cumulant", "cumular", "cumulate", "cumulating", "cumulation", "cumulatist", "cumulative", "cumulene", "cumulet", "cumuli", "cumulocirrus", "cumulonimbus", "cumulophyric", "cumulose", "cumulostratus", "cumulous", "cumulus", "cumyl", "curassow", "curcuma", "cushite", "cushitic", "cutwater", "cyanus", "cystitome", "dagassa", "daintith", "danalite", "dassy", "dawcock", "dawtit", "debarrass", "deboner", "debugger", "decimosexto", "decreptitude", "decuman", "decumbence", "decumbency", "decumbiture", "defensemen", "degass", "deglutitory", "deminude", "denude", "desex", "destituent", "destitute", "destituting", "destitution", "dhanush", "diaconicum", "dickcissel", "dickens", "dicker", "dickey", "dickie", "dickinsonite", "dickite", "dicksonia", "dickty", "dicky", "disexcommunicate", "disexercise", "distinctity", "document", "doronicum", "duniewassal", "duniwassal", "dyked", "dykehopper", "dyker", "earmuff", "ecumenacy", "ecumenic", "ecumenism", "ecumenist", "ecumenopolis", "efecks", "eightballs", "elanus", "elasticum", "elkoshite", "embarrass", "entitative", "entity", "epipubes", "epoophoron", "epornitic", "esexual", "essex", "ethanal", "exactitude", "exacum", "excisemen", "eyass", "eyeballs", "factitude", "fanal", "farse", "fass", "fastballs", "fecket", "feckful", "feckless", "feckly", "feckulence", "fellated", "fellatee", "fellation", "fireballs", "fittit", "flanged", "flangeless", "flanger", "flangeway", "flocoon", "floodcock", "footballs", "fortitude", "fortitudinous", "frass", "fusicoccum", "gaiassa", "galabia", "galeass", "galleass", "galliass", "gamecock", "garse", "gassy", "gearset", "gentianal", "ghassanid", "girgashite", "gittith", "goldtit", "goniatitoid", "goofballs", "gorcock", "graffage", "grass", "gratitude", "guaiacum", "guaiocum", "gyassa", "haboob", "hadassah", "hairballs", "halfcock", "hancockite", "handballs", "harass", "hardballs", "harshitha", "hassar", "hassle", "hassling", "hauberticum", "haycock", "hearse", "heartwater", "heelballs", "hemiganus", "hemipenis", "heptitol", "heterosex", "hexanal", "highballs", "hittitology", "hoarse", "homosexual", "honeyballs", "hornyhanded", "hornyhead", "horsemen", "horseshit", "hosecock", "housemen", "humbugger", "hypericum", "hyperprosexia", "hypersexual", "hypoprosexia", "ianus", "ifecks", "impuberate", "incestuous", "incumbant", "incumbence", "incumbencies", "incumbency", "incumbition", "ineptitude", "institor", "institue", "institute", "instituting", "institution", "institutive", "institutor", "institutress", "institutrix", "interflange", "intersex", "investitor", "irenicum", "ischioanal", "isophanal", "jactitate", "jactitating", "jactitation", "janus", "jass", "jedcock", "jerked", "jerker", "jerkier", "jerkies", "jerkily", "jerkin", "jerkish", "jerksome", "jerkwater", "jerky", "jitterbugger", "jiveass", "jizzen", "jocum", "judcock", "kaneshite", "kassabah", "kassak", "kassu", "katharses", "kavass", "khass", "khorassan", "koreishite", "kvass", "labial", "labiatae", "labiate", "labiatiflorous", "laconicum", "laocoon", "lapcock", "lass", "latitat", "latitude", "latitudinal", "latitudinarian", "latitudinary", "latitudinous", "leafage", "leasemen", "lentitude", "lentitudinous", "levisticum", "ligusticum", "liripoop", "lobcock", "locum", "logcock", "lubbercock", "lucanus", "lucuma", "lucumia", "lucumo", "mackintoshite", "macumba", "madagass", "madrassah", "maecenasship", "magnacumlaude", "makassar", "manal", "manificum", "manus", "marrymuffe", "marse", "marshite", "mass", "maycock", "meatballs", "mechanal", "mecum", "medick", "megalopenis", "megass", "meltith", "meltwater", "micropenis", "minahassa", "misexample", "misexecute", "misexecution", "misexpectation", "misexpend", "misexplain", "misexplanation", "misexplicate", "misexplication", "misexposition", "misexpound", "misexpress", "mishit", "mocock", "modicum", "monbuttu", "monosexualities", "monosexuality", "moorcock", "morass", "mothballs", "motherfucker", "motitation", "muffed", "muffer", "muffet", "muffin", "muffish", "muffle", "mufflin", "muffy", "muircock", "multitagged", "multitask", "multithread", "multitoed", "multitoned", "multitube", "multitubular", "multitude", "multitudinal", "multitudinary", "multitudinism", "multitudinist", "multitudinosity", "multitudinous", "multiturn", "nagkassar", "nakeder", "nakedest", "nakedish", "nakedize", "nakedly", "nakedness", "nakedweed", "nakedwood", "nassa", "nassology", "nictitate", "nictitating", "nictitation", "niddick", "niddicock", "niggard", "nincompoop", "nincum", "nocument", "noncock", "nonporness", "nonsexist", "nonsexlinked", "nonsexual", "norsemen", "nothofagus", "notidanus", "nyctitropic", "nyctitropism", "oarcock", "oceanus", "oddballs", "oecumenian", "oouassa", "organal", "orichalcum", "oroanal", "orthoganal", "ostracum", "outwatch", "outwater", "overfag", "overprick", "oversexed", "paddlecock", "paganalia", "pandanales", "pandanus", "panegyricum", "panicum", "pansexism", "pansexual", "panus", "parabolanus", "paratitla", "paratitlon", "parnassus", "parse", "partita", "partitura", "pass", "patacoon", "patchcock", "pelecanus", "penistone", "pentit", "perianal", "perqueer", "persephassa", "peshito", "petcock", "petit", "pettitoes", "philopornist", "piassaba", "piassava", "picumninae", "picumnus", "pillicock", "pinballs", "pinchcock", "pinprick", "pissabed", "pissant", "pissasphalt", "pissed", "pissing", "pissodes", "pissoir", "placuntoma", "platanus", "platitudinisation", "platitudinise", "platitudinising", "platitudinism", "platitudinist", "platitudinization", "platitudinize", "platitudinizing", "plotcock", "poetito", "polabian", "poortith", "poppycock", "pornerastic", "porno", "portitor", "possemen", "postanal", "posticum", "practicum", "praeanal", "preanal", "prickado", "prickant", "pricked", "pricker", "pricket", "prickfoot", "prickier", "prickiest", "pricking", "prickish", "prickle", "pricklier", "prickliest", "prickliness", "prickling", "pricklouse", "prickly", "prickmadam", "prickmedainty", "prickproof", "prickseam", "prickshot", "prickspur", "pricktimber", "prickwood", "pricky", "princock", "promptitude", "pseudolabia", "psychoanal", "psychosexual", "puberal", "pubertal", "pubertic", "puberties", "puberty", "puberulent", "puberulous", "pubescence", "pubescency", "pubescent", "publicum", "puccoon", "puffballs", "puirtith", "pushballs", "pussycat", "pussyfoot", "pussytoe", "quantitate", "quantitation", "quantitative", "quantity", "quass", "queered", "queerer", "queerest", "queering", "queerish", "queerity", "queerly", "queerness", "queersome", "queery", "quinisext", "quotity", "raccoon", "racoon", "ranal", "raphanus", "rassasy", "rassle", "rassling", "ratitae", "ratitous", "rebuttable", "rebuttably", "recock", "rectitude", "rectitudinous", "recumb", "rejerk", "remittitur", "resex", "restitue", "restitute", "restituting", "restitution", "restitutive", "restitutor", "ribboner", "robustity", "roofage", "routhercock", "rusticum", "saccoon", "sackbutt", "saltwater", "sanctitude", "sanctity", "sargasso", "sargassum", "sarsechim", "sarsen", "sass", "saturday", "scantity", "scoon", "scrap", "screwballs", "scriptitory", "scuddick", "sculptitory", "scum", "scuttlebutt", "semence", "semencinae", "semencontra", "sement", "semihorny", "seminaked", "seminude", "semisextile", "senectitude", "serfage", "serranus", "sesquisextal", "sesquitertianal", "sheafage", "shinnecock", "shirlcock", "shitepoke", "shitheel", "shither", "shredcock", "shuttlecock", "sicilicum", "siganus", "sillcock", "silvanus", "skiddycock", "skirlcock", "slutch", "sluther", "sluthood", "snaked", "snowballs", "sociosexual", "softballs", "solanal", "sortita", "sourballs", "sparassodont", "spiss", "spitballs", "spitchcock", "spunked", "spunkie", "spunkily", "spunkiness", "spunking", "spunkless", "spunky", "squawtits", "staffage", "stalactital", "stitch", "stith", "stituted", "stopcock", "stormcock", "stuffage", "sturdied", "sturdier", "sturdiest", "sturdily", "sturdy", "subnude", "subsextuple", "substituent", "substitutabilities", "substitutability", "substitutable", "substitute", "substituting", "substitution", "substitutive", "succumb", "suffraganal", "sulfaguanidine", "sulpharsenid", "supersex", "sussex", "swank", "sweetwater", "swordick", "sycock", "synarses", "tabacum", "tagassu", "talcum", "tanala", "tapisser", "tapissier", "taraxacum", "tarrass", "tarse", "tass", "tautit", "tayassu", "teanal", "tecum", "tenochtitlan", "terass", "tetanal", "tetanus", "thorny", "thricecock", "tillicum", "titan", "titar", "titbit", "tite", "titfer", "titfish", "tithable", "tithal", "tithe", "tithing", "tithonia", "tithonic", "tithonographic", "tithonometer", "tithonus", "tithymal", "titi", "titlark", "title", "titlike", "titling", "titlist", "titmal", "titman", "titmarsh", "titmen", "titmice", "titmmice", "titmouse", "titoism", "titoist", "titoki", "titrable", "titrant", "titratable", "titrate", "titrating", "titration", "titrator", "titre", "titrimetric", "titrimetry", "titter", "tittivate", "tittivating", "tittivation", "tittivator", "tittle", "tittlin", "tittup", "titty", "titubancy", "titubant", "titubate", "titubation", "titulado", "titular", "titulation", "titule", "tituli", "titulus", "titurel", "titus", "toddick", "tomtit", "totanus", "towcock", "toxicum", "trampcock", "transexperiental", "transexperiential", "transpenisular", "transsexual", "trapballs", "trass", "tripsacum", "triticum", "tucum", "turdetan", "turdidae", "turdiform", "turdinae", "turdine", "turdoid", "turdus", "turfage", "turncock", "twank", "twatchel", "twatterlight", "twattle", "twattling", "tycoon", "tympanal", "typicum", "unanalyzably", "unbiassable", "uncock", "uncrossexaminable", "uncrossexamined", "undercumstand", "undersexed", "undersexton", "unfagged", "unfagoted", "unhorny", "unicum", "unisex", "unnaked", "unporness", "unprostitute", "unsex", "untithability", "upcock", "upjerk", "upprick", "uranus", "urethrosexual", "urinosexual", "urucum", "vakass", "valerianales", "vandyke", "varanus", "vassal", "vassar", "vassos", "vastitude", "vastity", "vavassor", "vectitation", "versemen", "vestiture", "viaticum", "vitita", "volcanus", "volleyballs", "wambutti", "warse", "washita", "weathercock", "wereass", "wharfage", "whisterpoop", "whitehass", "winterdykes", "wistit", "woodcock", "woollybutt", "wristwatch", "yercum", "zaddick", "zimentwater", "zincum", "zoacum", "zygomaticum"}