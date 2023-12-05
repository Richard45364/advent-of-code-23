class Map:

    def __init__(self, mappings):
        self.mappings = []
        for input_map in mappings:
            map_from = (input_map[1], input_map[1] + input_map[2])
            map_to = (input_map[0], input_map[0] + input_map[2])
            self.mappings.append({"map_from": map_from, "map_to": map_to})
        self.mappings = sorted(self.mappings, key=lambda x: x['map_from'][0])

    def propagate(self, seed):
        for mapping in self.mappings:
            if mapping["map_from"][0] <= seed < mapping["map_from"][1]:
                return mapping["map_to"][0] + (seed-mapping["map_from"][0])
        return seed

    def get_mappings(self):
        return self.mappings


def main():
    seeds = [79, 14, 55, 13]

    seed_to_soil_map = Map([
        (50, 98, 2),
        (52, 50, 48),
    ])

    soil_to_fertilizer_map  = Map([
                      (0, 15, 37),
                      (37, 52, 2),
                      (39, 0, 15)
    ])

    fertilizer_to_water_map = Map([
        (49, 53, 8),
        (0, 11, 42),
        (42, 0, 7),
        (57, 7, 4)
    ])

    water_to_light_map = Map([
        (88, 18, 7),
        (18, 25, 70)
    ])

    light_to_temperature_map = Map([
        (45, 77, 23),
        (81, 45, 19),
        (68, 64, 13)
    ])

    temperature_to_humidity_map = Map([
        (0, 69, 1),
        (1, 0, 69)
    ])

    humidity_to_location_map = Map([
        (60, 56, 37),
        (56, 93, 4)
    ])

    for seed in seeds:
        result = humidity_to_location_map.propagate(
            temperature_to_humidity_map.propagate(
                light_to_temperature_map.propagate(
                    water_to_light_map.propagate(
                        fertilizer_to_water_map.propagate(
                            soil_to_fertilizer_map.propagate(
                                seed_to_soil_map.propagate(seed)
                            )
                        )
                    )
                )
            )
        )
        print(result)


if __name__ == "__main__":
    main()
