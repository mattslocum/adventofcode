import Foundation
import RegexBuilder

struct Data13 {
    // arrays are always X, Y
    var A: [Int]?
    var B: [Int]?
    var Prise: [Int]?
}

struct Challenge13: Challenge, Identifiable {
    var id = UUID()
    var day: String { "13" }
    var description: String { """
        Part 1: \n
        Part 2: 
        """ }
    var input: String
    
    init() {
        input = (try? readTextFile(file: "13Test")) ?? ""
    }

    private func parseInput() throws -> ([Data13]) {
        var machines: [Data13] = []
        let sets = input.components(separatedBy: "\n\n")
        for set in sets {
            var machine: Data13 = Data13()
            var lines = set.components(separatedBy: CharacterSet.newlines)
            let removeA = lines[0].range(of: "Button A: ")
            lines[0].removeSubrange(removeA!)
            let partsA = lines[0].components(separatedBy: ", ")
            
            let removeB = lines[1].range(of: "Button B: ")
            lines[1].removeSubrange(removeB!)
            let removeP = lines[2].range(of: "Prize: ")
            lines[2].removeSubrange(removeP!)
        }
//        for line in lines {
//            var points: [String] = []
//            for val in line {
//                points.append(String(val))
//            }
//            map.append(points)
//        }
        return machines
    }

    func part1() async throws -> String {
        let machines = try parseInput()
        var total = 0

        return String(total)
    }
    
    func part2() async throws -> String {
        
        return "part 2"
    }
}



