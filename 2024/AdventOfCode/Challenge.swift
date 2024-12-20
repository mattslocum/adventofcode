import Foundation

protocol Challenge: Identifiable {
    var id: UUID { get }
    var day: String { get }
    var description: String { get }
    var input: String { get }

    func part1() async throws -> String
    func part2() async throws -> String
}

class ChallengeLoader {
    static func loadChallenges() -> [any Challenge] {
        return [
            Challenge01(),
            Challenge02(),
            Challenge03(),
            Challenge04(),
            Challenge05(),
            Challenge06(),
            Challenge07(),
            Challenge08(),
            Challenge09(),
            Challenge10(),
            Challenge11(),
            Challenge12(),
            Challenge13(),
            // Add more challenges here
        ]
    }
}
