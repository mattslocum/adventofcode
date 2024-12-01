import Foundation

protocol Challenge: Identifiable {
    var id: UUID { get }
    var day: String { get }

    func run1() async throws -> String
    func run2() async throws -> String
}

class ChallengeLoader {
    static func loadChallenges() -> [any Challenge] {
        return [
            Challenge01(),
            Challenge02(),
            // Add more challenges here
        ]
    }
}
