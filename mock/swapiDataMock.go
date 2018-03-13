package mock

//Planet1 mock data
const Planet1 string = `{
    "name": "Tatooine",
    "rotation_period": "23",
    "orbital_period": "304",
    "diameter": "10465",
    "climate": "arid",
    "gravity": "1 standard",
    "terrain": "desert",
    "surface_water": "1",
    "population": "200000",
    "residents": [
        "https://swapi.co/api/people/1/",
        "https://swapi.co/api/people/2/",
        "https://swapi.co/api/people/4/",
        "https://swapi.co/api/people/6/",
        "https://swapi.co/api/people/7/",
        "https://swapi.co/api/people/8/",
        "https://swapi.co/api/people/9/",
        "https://swapi.co/api/people/11/",
        "https://swapi.co/api/people/43/",
        "https://swapi.co/api/people/62/"
    ],
    "films": [
        "https://swapi.co/api/films/5/",
        "https://swapi.co/api/films/4/",
        "https://swapi.co/api/films/6/",
        "https://swapi.co/api/films/3/",
        "https://swapi.co/api/films/1/"
    ],
    "created": "2014-12-09T13:50:49.641000Z",
    "edited": "2014-12-21T20:48:04.175778Z",
    "url": "https://swapi.co/api/planets/1/"
}`

//Planet2 mock data
const Planet2 string = `{
    "name": "Chandrila",
    "rotation_period": "20",
    "orbital_period": "368",
    "diameter": "13500",
    "climate": "temperate",
    "gravity": "1",
    "terrain": "plains, forests",
    "surface_water": "40",
    "population": "1200000000",
    "residents": [
        "https://swapi.co/api/people/28/"
    ],
    "films": [],
    "created": "2014-12-18T11:11:51.872000Z",
    "edited": "2014-12-20T20:58:18.472000Z",
    "url": "https://swapi.co/api/planets/32/"
}`

//PlanetPage1 mock data
const PlanetPage1 string = `{
    "count": 21,
    "next": "https://swapi.co/api/planets/?page=2",
    "previous": null,
    "results": [
        {
            "name": "Alderaan",
            "rotation_period": "24",
            "orbital_period": "364",
            "diameter": "12500",
            "climate": "temperate",
            "gravity": "1 standard",
            "terrain": "grasslands, mountains",
            "surface_water": "40",
            "population": "2000000000",
            "residents": [
                "https://swapi.co/api/people/5/",
                "https://swapi.co/api/people/68/",
                "https://swapi.co/api/people/81/"
            ],
            "films": [
                "https://swapi.co/api/films/6/",
                "https://swapi.co/api/films/1/"
            ],
            "created": "2014-12-10T11:35:48.479000Z",
            "edited": "2014-12-20T20:58:18.420000Z",
            "url": "https://swapi.co/api/planets/2/"
        },
        {
            "name": "Yavin IV",
            "rotation_period": "24",
            "orbital_period": "4818",
            "diameter": "10200",
            "climate": "temperate, tropical",
            "gravity": "1 standard",
            "terrain": "jungle, rainforests",
            "surface_water": "8",
            "population": "1000",
            "residents": [],
            "films": [
                "https://swapi.co/api/films/1/"
            ],
            "created": "2014-12-10T11:37:19.144000Z",
            "edited": "2014-12-20T20:58:18.421000Z",
            "url": "https://swapi.co/api/planets/3/"
        },
        {
            "name": "Hoth",
            "rotation_period": "23",
            "orbital_period": "549",
            "diameter": "7200",
            "climate": "frozen",
            "gravity": "1.1 standard",
            "terrain": "tundra, ice caves, mountain ranges",
            "surface_water": "100",
            "population": "unknown",
            "residents": [],
            "films": [
                "https://swapi.co/api/films/2/"
            ],
            "created": "2014-12-10T11:39:13.934000Z",
            "edited": "2014-12-20T20:58:18.423000Z",
            "url": "https://swapi.co/api/planets/4/"
        },
        {
            "name": "Dagobah",
            "rotation_period": "23",
            "orbital_period": "341",
            "diameter": "8900",
            "climate": "murky",
            "gravity": "N/A",
            "terrain": "swamp, jungles",
            "surface_water": "8",
            "population": "unknown",
            "residents": [],
            "films": [
                "https://swapi.co/api/films/2/",
                "https://swapi.co/api/films/6/",
                "https://swapi.co/api/films/3/"
            ],
            "created": "2014-12-10T11:42:22.590000Z",
            "edited": "2014-12-20T20:58:18.425000Z",
            "url": "https://swapi.co/api/planets/5/"
        },
        {
            "name": "Bespin",
            "rotation_period": "12",
            "orbital_period": "5110",
            "diameter": "118000",
            "climate": "temperate",
            "gravity": "1.5 (surface), 1 standard (Cloud City)",
            "terrain": "gas giant",
            "surface_water": "0",
            "population": "6000000",
            "residents": [
                "https://swapi.co/api/people/26/"
            ],
            "films": [
                "https://swapi.co/api/films/2/"
            ],
            "created": "2014-12-10T11:43:55.240000Z",
            "edited": "2014-12-20T20:58:18.427000Z",
            "url": "https://swapi.co/api/planets/6/"
        },
        {
            "name": "Endor",
            "rotation_period": "18",
            "orbital_period": "402",
            "diameter": "4900",
            "climate": "temperate",
            "gravity": "0.85 standard",
            "terrain": "forests, mountains, lakes",
            "surface_water": "8",
            "population": "30000000",
            "residents": [
                "https://swapi.co/api/people/30/"
            ],
            "films": [
                "https://swapi.co/api/films/3/"
            ],
            "created": "2014-12-10T11:50:29.349000Z",
            "edited": "2014-12-20T20:58:18.429000Z",
            "url": "https://swapi.co/api/planets/7/"
        },
        {
            "name": "Naboo",
            "rotation_period": "26",
            "orbital_period": "312",
            "diameter": "12120",
            "climate": "temperate",
            "gravity": "1 standard",
            "terrain": "grassy hills, swamps, forests, mountains",
            "surface_water": "12",
            "population": "4500000000",
            "residents": [
                "https://swapi.co/api/people/3/",
                "https://swapi.co/api/people/21/",
                "https://swapi.co/api/people/36/",
                "https://swapi.co/api/people/37/",
                "https://swapi.co/api/people/38/",
                "https://swapi.co/api/people/39/",
                "https://swapi.co/api/people/42/",
                "https://swapi.co/api/people/60/",
                "https://swapi.co/api/people/61/",
                "https://swapi.co/api/people/66/",
                "https://swapi.co/api/people/35/"
            ],
            "films": [
                "https://swapi.co/api/films/5/",
                "https://swapi.co/api/films/4/",
                "https://swapi.co/api/films/6/",
                "https://swapi.co/api/films/3/"
            ],
            "created": "2014-12-10T11:52:31.066000Z",
            "edited": "2014-12-20T20:58:18.430000Z",
            "url": "https://swapi.co/api/planets/8/"
        },
        {
            "name": "Coruscant",
            "rotation_period": "24",
            "orbital_period": "368",
            "diameter": "12240",
            "climate": "temperate",
            "gravity": "1 standard",
            "terrain": "cityscape, mountains",
            "surface_water": "unknown",
            "population": "1000000000000",
            "residents": [
                "https://swapi.co/api/people/34/",
                "https://swapi.co/api/people/55/",
                "https://swapi.co/api/people/74/"
            ],
            "films": [
                "https://swapi.co/api/films/5/",
                "https://swapi.co/api/films/4/",
                "https://swapi.co/api/films/6/",
                "https://swapi.co/api/films/3/"
            ],
            "created": "2014-12-10T11:54:13.921000Z",
            "edited": "2014-12-20T20:58:18.432000Z",
            "url": "https://swapi.co/api/planets/9/"
        },
        {
            "name": "Kamino",
            "rotation_period": "27",
            "orbital_period": "463",
            "diameter": "19720",
            "climate": "temperate",
            "gravity": "1 standard",
            "terrain": "ocean",
            "surface_water": "100",
            "population": "1000000000",
            "residents": [
                "https://swapi.co/api/people/22/",
                "https://swapi.co/api/people/72/",
                "https://swapi.co/api/people/73/"
            ],
            "films": [
                "https://swapi.co/api/films/5/"
            ],
            "created": "2014-12-10T12:45:06.577000Z",
            "edited": "2014-12-20T20:58:18.434000Z",
            "url": "https://swapi.co/api/planets/10/"
        },
        {
            "name": "Geonosis",
            "rotation_period": "30",
            "orbital_period": "256",
            "diameter": "11370",
            "climate": "temperate, arid",
            "gravity": "0.9 standard",
            "terrain": "rock, desert, mountain, barren",
            "surface_water": "5",
            "population": "100000000000",
            "residents": [
                "https://swapi.co/api/people/63/"
            ],
            "films": [
                "https://swapi.co/api/films/5/"
            ],
            "created": "2014-12-10T12:47:22.350000Z",
            "edited": "2014-12-20T20:58:18.437000Z",
            "url": "https://swapi.co/api/planets/11/"
        }
    ]
}`

//PlanetPage2 mock data
const PlanetPage2 string = `{
    "count": 21,
    "next": "https://swapi.co/api/planets/?page=3",
    "previous": "https://swapi.co/api/planets/?page=2",
    "results": [
        {
            "name": "Chandrila",
            "rotation_period": "20",
            "orbital_period": "368",
            "diameter": "13500",
            "climate": "temperate",
            "gravity": "1",
            "terrain": "plains, forests",
            "surface_water": "40",
            "population": "1200000000",
            "residents": [
                "https://swapi.co/api/people/28/"
            ],
            "films": [],
            "created": "2014-12-18T11:11:51.872000Z",
            "edited": "2014-12-20T20:58:18.472000Z",
            "url": "https://swapi.co/api/planets/32/"
        },
        {
            "name": "Sullust",
            "rotation_period": "20",
            "orbital_period": "263",
            "diameter": "12780",
            "climate": "superheated",
            "gravity": "1",
            "terrain": "mountains, volcanoes, rocky deserts",
            "surface_water": "5",
            "population": "18500000000",
            "residents": [
                "https://swapi.co/api/people/31/"
            ],
            "films": [],
            "created": "2014-12-18T11:25:40.243000Z",
            "edited": "2014-12-20T20:58:18.474000Z",
            "url": "https://swapi.co/api/planets/33/"
        },
        {
            "name": "Toydaria",
            "rotation_period": "21",
            "orbital_period": "184",
            "diameter": "7900",
            "climate": "temperate",
            "gravity": "1",
            "terrain": "swamps, lakes",
            "surface_water": "unknown",
            "population": "11000000",
            "residents": [
                "https://swapi.co/api/people/40/"
            ],
            "films": [],
            "created": "2014-12-19T17:47:54.403000Z",
            "edited": "2014-12-20T20:58:18.476000Z",
            "url": "https://swapi.co/api/planets/34/"
        },
        {
            "name": "Malastare",
            "rotation_period": "26",
            "orbital_period": "201",
            "diameter": "18880",
            "climate": "arid, temperate, tropical",
            "gravity": "1.56",
            "terrain": "swamps, deserts, jungles, mountains",
            "surface_water": "unknown",
            "population": "2000000000",
            "residents": [
                "https://swapi.co/api/people/41/"
            ],
            "films": [],
            "created": "2014-12-19T17:52:13.106000Z",
            "edited": "2014-12-20T20:58:18.478000Z",
            "url": "https://swapi.co/api/planets/35/"
        },
        {
            "name": "Dathomir",
            "rotation_period": "24",
            "orbital_period": "491",
            "diameter": "10480",
            "climate": "temperate",
            "gravity": "0.9",
            "terrain": "forests, deserts, savannas",
            "surface_water": "unknown",
            "population": "5200",
            "residents": [
                "https://swapi.co/api/people/44/"
            ],
            "films": [],
            "created": "2014-12-19T18:00:40.142000Z",
            "edited": "2014-12-20T20:58:18.480000Z",
            "url": "https://swapi.co/api/planets/36/"
        },
        {
            "name": "Ryloth",
            "rotation_period": "30",
            "orbital_period": "305",
            "diameter": "10600",
            "climate": "temperate, arid, subartic",
            "gravity": "1",
            "terrain": "mountains, valleys, deserts, tundra",
            "surface_water": "5",
            "population": "1500000000",
            "residents": [
                "https://swapi.co/api/people/45/",
                "https://swapi.co/api/people/46/"
            ],
            "films": [],
            "created": "2014-12-20T09:46:25.740000Z",
            "edited": "2014-12-20T20:58:18.481000Z",
            "url": "https://swapi.co/api/planets/37/"
        },
        {
            "name": "Aleen Minor",
            "rotation_period": "unknown",
            "orbital_period": "unknown",
            "diameter": "unknown",
            "climate": "unknown",
            "gravity": "unknown",
            "terrain": "unknown",
            "surface_water": "unknown",
            "population": "unknown",
            "residents": [
                "https://swapi.co/api/people/47/"
            ],
            "films": [],
            "created": "2014-12-20T09:52:23.452000Z",
            "edited": "2014-12-20T20:58:18.483000Z",
            "url": "https://swapi.co/api/planets/38/"
        },
        {
            "name": "Vulpter",
            "rotation_period": "22",
            "orbital_period": "391",
            "diameter": "14900",
            "climate": "temperate, artic",
            "gravity": "1",
            "terrain": "urban, barren",
            "surface_water": "unknown",
            "population": "421000000",
            "residents": [
                "https://swapi.co/api/people/48/"
            ],
            "films": [],
            "created": "2014-12-20T09:56:58.874000Z",
            "edited": "2014-12-20T20:58:18.485000Z",
            "url": "https://swapi.co/api/planets/39/"
        },
        {
            "name": "Troiken",
            "rotation_period": "unknown",
            "orbital_period": "unknown",
            "diameter": "unknown",
            "climate": "unknown",
            "gravity": "unknown",
            "terrain": "desert, tundra, rainforests, mountains",
            "surface_water": "unknown",
            "population": "unknown",
            "residents": [
                "https://swapi.co/api/people/49/"
            ],
            "films": [],
            "created": "2014-12-20T10:01:37.395000Z",
            "edited": "2014-12-20T20:58:18.487000Z",
            "url": "https://swapi.co/api/planets/40/"
        },
        {
            "name": "Tund",
            "rotation_period": "48",
            "orbital_period": "1770",
            "diameter": "12190",
            "climate": "unknown",
            "gravity": "unknown",
            "terrain": "barren, ash",
            "surface_water": "unknown",
            "population": "0",
            "residents": [
                "https://swapi.co/api/people/50/"
            ],
            "films": [],
            "created": "2014-12-20T10:07:29.578000Z",
            "edited": "2014-12-20T20:58:18.489000Z",
            "url": "https://swapi.co/api/planets/41/"
        }
    ]
}`

//PlanetPage3 mock data
const PlanetPage3 string = `{
    "count": 21,
    "next": null,
    "previous": "https://swapi.co/api/planets/?page=2",
    "results": [
        {
            "name": "Jakku",
            "rotation_period": "unknown",
            "orbital_period": "unknown",
            "diameter": "unknown",
            "climate": "unknown",
            "gravity": "unknown",
            "terrain": "deserts",
            "surface_water": "unknown",
            "population": "unknown",
            "residents": [],
            "films": [
                "https://swapi.co/api/films/7/"
            ],
            "created": "2015-04-17T06:55:57.556495Z",
            "edited": "2015-04-17T06:55:57.556551Z",
            "url": "https://swapi.co/api/planets/61/"
        }
    ]
}`
