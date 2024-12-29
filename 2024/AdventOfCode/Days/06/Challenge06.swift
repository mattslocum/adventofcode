import Foundation
import RegexBuilder

enum Type06 {
    case Empty
    case Visited
    case Blocked
}

struct Challenge06: Challenge, Identifiable {
    var id = UUID()
    var day: String { "06" }
    var description: String { """
        Part 1: The cursor moves until it hits an obsticle and then turns right until is leaves the grid. Count the number of visited squares.\n
        Part 2: Find how many loops can be made by adding a single blocker.
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "06")) ?? ""
    }

    private func parseInput() throws -> ([[Type06]], [Int], Direction) {
        var map: [[Type06]] = []
        var pos = [0, 0]
        let dir = Direction.up
        
        let lines = input.components(separatedBy: CharacterSet.newlines)
        for (y, line) in lines.enumerated() {
            if line == "" {
                continue
            }
            map.append([])
            for (x, char) in line.enumerated() {
                if char == "#" {
                    map[y].append(.Blocked)
                } else if char == "^" {
                    map[y].append(.Visited)
                    pos = [y, x]
                } else {
                    map[y].append(.Empty)
                }
            }
        }
        return (map, pos, dir)
    }

    func part1() async throws -> String {
        var (map, pos, dir) = try parseInput()
        var isValid = true
        
        repeat {
            (map, pos, dir, isValid) = step(map:map, pos:pos, dir:dir)
        } while (isValid)

        var found = 0
        for row in map {
            for cell in row {
                if cell == .Visited {
                    found += 1
                }
            }
        }

        return String(found)
    }
    
    func step(map: [[Type06]], pos: [Int], dir: Direction) -> ([[Type06]], [Int], Direction, Bool) {
        var nextPos = pos
        var dir = dir
        var map = map
        
        switch dir {
        case .up:
            nextPos[0] -= 1
        case .down:
            nextPos[0] += 1
        case .left:
            nextPos[1] -= 1
        case .right:
            nextPos[1] += 1
        }
        
        // did we run off the grid?
        if nextPos[0] < 0 || nextPos[0] >= map.count || nextPos[1] < 0 || nextPos[1] >= map[0].count {
            return (map, nextPos, dir, false)
        }
        
        if map[nextPos[0]][nextPos[1]] == .Blocked {
            switch dir {
            case .up:
                dir = .right
            case .right:
                dir = .down
            case .down:
                dir = .left
            case .left:
                dir = .up
            }
            nextPos = pos
        } else {
            map[nextPos[0]][nextPos[1]] = .Visited
        }
        
        return (map, nextPos, dir, true)
    }

    // 1523
    func part2() async throws -> String {
        let (map, pos, dir) = try parseInput()

        var map2 = map
        var pos2 = pos
        var dir2 = dir
        var isValid = true
        var found: [[Int]] = []
        var blockPos = pos
        repeat {
            blockPos = pos2
            switch dir2 {
            case .up:
                if pos2[0] != 0 {
                    blockPos[0] -= 1
                }
            case .right:
                if pos2[1] != map2[0].count-1 {
                    blockPos[1] += 1
                }
            case .down:
                if pos2[0] != map2.count-1 {
                    blockPos[0] += 1
                }
            case .left:
                if pos2[1] != 0 {
                    blockPos[1] -= 1
                }
            }
            map2[blockPos[0]][blockPos[1]] = .Blocked
            
            // check loop from origin start and dir, but new map
            if !found.contains(blockPos) && isLoop(map: map2, pos: pos, dir: dir) {
                found.append(blockPos)
//                print("found loop at \(blockPos)")
            }
//            printGrid(map: map2)

            // send old map so we don't block our path
            (map2, pos2, dir2, isValid) = step(map:map, pos:pos2, dir:dir2)
        } while (isValid)

        return String(found.count)
    }

    func isLoop(map: [[Type06]], pos: [Int], dir: Direction) -> Bool {
        var isValid = true
        var pos = pos
        var dir = dir
        var newDir = dir
        var newPos = pos
        var collision: [[Int]: [Direction]] = [:]
        
        repeat {
            (_, newPos, newDir, isValid) = step(map:map, pos:pos, dir:dir)
            // check for collisions, and track previsous positions
            if dir != newDir {
                // check if we've been here before in this direction
                if collision[newPos, default: []].contains(newDir) {
                    return true
                }
                collision[newPos, default: []].append(newDir)
            }
            pos = newPos
            dir = newDir
        } while (isValid)
        
        return isValid
    }
    
    func printGrid(map: [[Type06]]) {
        print("Grid:")
        for row in map {
            var line = ""
            for cell in row {
                switch cell {
                case .Empty:
                    line += "."
                case .Visited:
                    line += "*"
                case .Blocked:
                    line += "#"
                }
            }
            print(line)
        }
    }
}



