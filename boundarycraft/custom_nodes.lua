minetest.register_node(
	"boundarycraft:forcefield",
	{
		description = "forcefield_glass",
		inventory_image = minetest.inventorycube("moreblocks_clean_glass.png"),
		drawtype = "glasslike",
		walkable = false,
		use_texture_alpha = true,
		paramtype = "light",
		sunlight_propagates = true,
		light_source = 10,
		pointable = false,
		post_effect_color = {r = 8, g = 225, b = 217, a = 50},
		tiles = {
			{
				name = "teleporter_glow_animated.png",
				animation = {type = "vertical_frames", aspect_w = 16, aspect_h = 16, length = 5.0}
			}
		},
		is_ground_content = true,
		groups = {unbreakable = 1}
	}
)
