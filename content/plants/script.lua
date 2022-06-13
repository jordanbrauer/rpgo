math.randomseed(os.time())

function SpawnTree(x, y, z)
    local season = "summer"
    local style = math.random(1, 8)

    return spawn(
        string.format("content/plants/models/trees/tree0%d.obj", style),
        string.format("content/plants/textures/trees/tree0%d_%s.png", style, season),
        x, y, z,
        2.0
    )
end

print(string.format("Spawned Tree: %d", SpawnTree(0.0, 0.0, 0.0)))
