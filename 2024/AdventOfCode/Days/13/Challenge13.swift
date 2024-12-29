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
        input = (try? readTextFile(file: "13")) ?? ""
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
        let start = min(m.Prise[0] / m.B[0], m.Prise[1] / m.B[1])
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
            total += findCostEquations(m: m)
        }

        return String(total)
    }
    
    func findCostEquations(m: Data13) -> Int {
        // Start with B because it is cheeper
        let axPrimes = primeFactors(n: m.A[0])
        let ayPrimes = primeFactors(n: m.A[1])
        let primes = maxMaps(a: axPrimes, b: ayPrimes)
        let target = buildNum(primes: primes)
        print("\(m.A[0]) and \(m.A[1]) Max Primes: \(primes) Target: \(target)")
        
        // make both equations have the same A value so they will cancel during subtraction
        var data = m
        var mult = target / m.A[0]
        data.B[0] *= mult
        data.Prise[0] *= mult
        mult = target / m.A[1]
        data.B[1] *= mult
        data.Prise[1] *= mult
        
        // subtract the equations to get B
        let bDiff = data.B[0] - data.B[1]
        let pDiff = data.Prise[0] - data.Prise[1]
        let b = pDiff / bDiff
        
        // solve for A
        let a = (m.Prise[0] - (m.B[0] * b)) / m.A[0]
        
        // We might not have a perfect match, so double check.
        if ((m.A[0] * a + m.B[0] * b) == m.Prise[0]) && ((m.A[1] * a + m.B[1] * b) == m.Prise[1]) {
            print("FOUND MATCH: \(a), \(b)")
            return (a * 3) + b
        }
//        print("FAILED")

        return 0
    }
    
    func primeFactors(n: Int) -> [Int: Int] {
        var num = n
        var factors : [Int: Int] = [:]
        while num > 1 {
            if num % 2 == 0 {
                factors[2, default: 0] += 1
                num = num / 2
                continue
            }
            // TODO: be more efficient on increments of primes
            var n = 3
            while n <= num {
                if num % n == 0 {
                    factors[n, default: 0] += 1
                    num = num / n
                    break
                }
                n += 2
            }
        }
        return factors
    }
    
    func maxMaps(a: [Int: Int], b: [Int: Int]) -> [Int: Int] {
        var merged = a
        for (key, value) in b {
            if merged[key, default: 0] < value {
                merged[key] = value
            }
        }
        return merged
    }
    
    func buildNum(primes: [Int: Int]) -> Int {
        var num = 1
        for (key, value) in primes {
            num *= Int(pow(Double(key), Double(value)))
        }
        return num
    }
}



