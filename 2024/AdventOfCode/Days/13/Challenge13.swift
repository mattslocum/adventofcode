import Foundation
import RegexBuilder

struct Data13 {
    // arrays are always X, Y
    var A: [Int]
    var B: [Int]
    var Prise: [Int]
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

    private func parseInput(adj: Int) throws -> ([Data13]) {
        var machines: [Data13] = []
        let sets = input.components(separatedBy: "\n\n")
        for set in sets {
            var machine: Data13 = Data13(A: [], B: [], Prise: [])
            var lines = set.components(separatedBy: CharacterSet.newlines)

            let removeA = lines[0].range(of: "Button A: ")
            lines[0].removeSubrange(removeA!)
            var partsA = lines[0].components(separatedBy: ", ")
            var removeChars = partsA[0].range(of: "X+")
            partsA[0].removeSubrange(removeChars!)
            removeChars = partsA[1].range(of: "Y+")
            partsA[1].removeSubrange(removeChars!)
            machine.A = partsA.map({ Int($0)! })

            let removeB = lines[1].range(of: "Button B: ")
            lines[1].removeSubrange(removeB!)
            var partsB = lines[1].components(separatedBy: ", ")
            removeChars = partsB[0].range(of: "X+")
            partsB[0].removeSubrange(removeChars!)
            removeChars = partsB[1].range(of: "Y+")
            partsB[1].removeSubrange(removeChars!)
            machine.B = partsB.map({ Int($0)! })

            let removeP = lines[2].range(of: "Prize: ")
            lines[2].removeSubrange(removeP!)
            var partsP = lines[2].components(separatedBy: ", ")
            removeChars = partsP[0].range(of: "X=")
            partsP[0].removeSubrange(removeChars!)
            removeChars = partsP[1].range(of: "Y=")
            partsP[1].removeSubrange(removeChars!)
            machine.Prise = partsP.map({ Int($0)! + adj })

            machines.append(machine)
        }
        return machines
    }

    func part1() async throws -> String {
        let machines = try parseInput(adj: 0)
        var total = 0
        
        for m in machines {
            total += findCost(m: m)
        }

        return String(total)
    }
    
    func findCost(m: Data13) -> Int {
        // Start with B because it is cheeper
        let start = m.Prise[0] / m.B[0]
        // loop backwards from start to 0
        for b in (0..<start+1).reversed() {
            let posX = m.B[0] * b
            let remainderX = m.Prise[0] - posX
            if remainderX % m.A[0] == 0 {
                let a = remainderX / m.A[0]
//                print("FOUND MATCH X: \(a), \(b)")
                // we found an match on Xs. Now check Ys
//                print(m.B[1] * b, m.A[1] * a)
                if (m.B[1] * b + m.A[1] * a) == m.Prise[1] {
//                    print("FOUND MATCH: \(a), \(b)")
                    return (a * 3) + b
                }
            }
        }
        return 0
    }
    
    func part2() async throws -> String {
        let machines = try parseInput(adj: 10000000000000)
        var total = 0
        
        for m in machines {
            total += findCost(m: m)
        }

        return String(total)
    }
}



