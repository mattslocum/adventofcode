import Foundation
import RegexBuilder

struct Challenge12: Challenge, Identifiable {
    var id = UUID()
    var day: String { "12" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "12")) ?? ""
    }

    private func parseInput() throws -> ([[String]]) {
        var map: [[String]] = []
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for line in lines {
            var points: [String] = []
            for val in line {
                points.append(String(val))
            }
            map.append(points)
        }
        return map
    }

    // 1461752
    func part1() async throws -> String {
        let map = try parseInput()
        var checkedYX: [[Int]] = []
        var total = 0
        
        for y in 0..<map.count {
            for x in 0..<map[0].count {
                if !checkedYX.contains([y, x]) {
                    let (cells, perim) = findCellsAndPerimiter(map: map, startYX: [y, x])
//                    print(map[y][x], cells, perim)
                    checkedYX += cells
                    total += cells.count * perim
                }
            }
        }

        return String(total)
    }
    
    func findCellsAndPerimiter(map: [[String]], startYX: [Int]) -> ([[Int]], Int) {
        var list = [startYX]
        var perim = 0
        let findStr = map[startYX[0]][startYX[1]]
        
        var foundYX = [startYX]
        // start tree walking. Depth first search
        while foundYX.count > 0 {
            let curYX = foundYX.removeFirst()
            var sameNeighbor = 0;
            // check neighbors
            // up
            if curYX[0] > 0 && map[curYX[0]-1][curYX[1]] == findStr {
                let YX = [curYX[0]-1, curYX[1]]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                sameNeighbor += 1
            }
            // down
            if curYX[0] + 1 < map.count && map[curYX[0]+1][curYX[1]] == findStr {
                let YX = [curYX[0]+1, curYX[1]]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                sameNeighbor += 1
            }
            // left
            if curYX[1] > 0 && map[curYX[0]][curYX[1]-1] == findStr {
                let YX = [curYX[0], curYX[1]-1]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                sameNeighbor += 1
            }
            // right
            if curYX[1] + 1 < map[0].count && map[curYX[0]][curYX[1]+1] == findStr {
                let YX = [curYX[0], curYX[1]+1]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                sameNeighbor += 1
            }
            perim += 4 - sameNeighbor
        }

        return (list, perim)
    }

    func part2() async throws -> String {
        let map = try parseInput()
        var checkedYX: [[Int]] = []
        var total = 0
        
        for y in 0..<map.count {
            for x in 0..<map[0].count {
                if !checkedYX.contains([y, x]) {
                    let (cells, edges) = findCellsAndEdges(map: map, startYX: [y, x])
                    print(map[y][x], edges, cells.count * edges)
                    checkedYX += cells
                    total += cells.count * edges
                }
            }
        }

        return String(total)
    }
    
    // Note: edges are the same as corners
    func findCellsAndEdges(map: [[String]], startYX: [Int]) -> ([[Int]], Int) {
        var list = [startYX]
        var corner = 0
        let findStr = map[startYX[0]][startYX[1]]
        
        var foundYX = [startYX]
        // start tree walking. Depth first search
        while foundYX.count > 0 {
            let curYX = foundYX.removeFirst()
            var hasUp = false;
            var hasDown = false;
            var hasRight = false;
            var hasLeft = false;
            // check neighbors
            // up
            if curYX[0] > 0 && map[curYX[0]-1][curYX[1]] == findStr {
                let YX = [curYX[0]-1, curYX[1]]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                hasUp = true
            }
            // down
            if curYX[0] + 1 < map.count && map[curYX[0]+1][curYX[1]] == findStr {
                let YX = [curYX[0]+1, curYX[1]]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                hasDown = true
            }
            // left
            if curYX[1] > 0 && map[curYX[0]][curYX[1]-1] == findStr {
                let YX = [curYX[0], curYX[1]-1]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                hasLeft = true
            }
            // right
            if curYX[1] + 1 < map[0].count && map[curYX[0]][curYX[1]+1] == findStr {
                let YX = [curYX[0], curYX[1]+1]
                if !list.contains(YX) {
                    foundYX.append(YX)
                    list.append(YX)
                }
                hasRight = true
            }
            // check for outer corners
            if !hasUp && !hasRight {
                corner += 1
            }
            if !hasUp && !hasLeft {
                corner += 1
            }
            if !hasDown && !hasRight {
                corner += 1
            }
            if !hasDown && !hasLeft {
                corner += 1
            }
            // check for inner corners
            if hasUp && hasRight && map[curYX[0]-1][curYX[1]+1] != findStr {
                corner += 1
            }
            if hasUp && hasLeft && map[curYX[0]-1][curYX[1]-1] != findStr {
                corner += 1
            }
            if hasDown && hasRight && map[curYX[0]+1][curYX[1]+1] != findStr {
                corner += 1
            }
            if hasDown && hasLeft && map[curYX[0]+1][curYX[1]-1] != findStr {
                corner += 1
            }
        }

        return (list, corner)
    }
}



