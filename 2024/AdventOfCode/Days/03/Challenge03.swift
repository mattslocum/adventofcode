import Foundation
import RegexBuilder


let firstRef = Reference(Int.self)
let secondRef = Reference(Int.self)
let mulSearch = Regex {
    "mul("
    TryCapture(as: firstRef) {
        OneOrMore(.digit)
    } transform: { match in
        Int(match)
    }
    ","
    TryCapture(as: secondRef) {
        OneOrMore(.digit)
    } transform: { match in
        Int(match)
    }
    ")"
}
let doSearch = Regex {
    "do()"
}
let dontSearch = Regex {
    "don't()"
}

struct Challenge03: Challenge, Identifiable {
    var id = UUID()
    var day: String { "03" }
    var description: String { """
        Part 1: find all mul(a,b) and add all outcomes.\n
        Part 2: same as 1, but do() and don't() can be used to disable/enabled mul().
        """ }
    var input: String

    init() {
        input = (try? readTextFile(file: "03")) ?? ""
    }
    
    private func parseInput() throws -> (String) {
        return input
    }

    func part1() async throws -> String {
        let line = try parseInput()
        var answer = 0

        let muls = line.matches(of: mulSearch)
        for mul in muls {
            answer += mul[firstRef] * mul[secondRef]
        }

        return String(answer)
    }

    func part2() async throws -> String {
        var line = try parseInput()
        var answer = 0

        while let range = line.firstRange(of: mulSearch) {
            let nextDont = line.firstMatch(of: dontSearch)
            if nextDont != nil && nextDont!.startIndex < range.lowerBound {
                let nextDo = line.firstMatch(of: doSearch)
                let dontSubString = line[..<nextDo!.range.upperBound]
                line = line.replacingOccurrences(of: dontSubString, with: "")
                continue;
            }

            let mul = line.firstMatch(of: mulSearch)
            answer += mul![firstRef] * mul![secondRef]
            line.removeSubrange(range)
        }

        return String(answer)
    }
}



