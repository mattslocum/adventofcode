//
//  ContentView.swift
//  AdventOfCode
//
//  Created by Matt Slocum on 11/28/24.
//

import SwiftUI
import SwiftData

struct ContentView: View {
    @Environment(\.modelContext) private var modelContext
    @State private var days: [any Challenge] = ChallengeLoader.loadChallenges()

    var body: some View {
        NavigationSplitView {
            List {
                ForEach(days.indices, id: \.self) { index in
                    let day = days[index]
                    NavigationLink {
                        ChallengeDetailView(challenge: day)
                    } label: {
                        Text(day.day)
                    }
                }
            }
            .navigationSplitViewColumnWidth(min: 180, ideal: 200)
        } detail: {
            Text("Select a challenge")
        }
    }
}

#Preview {
    ContentView()
        .modelContainer(for: Item.self, inMemory: true)
}
