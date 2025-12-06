//
//  ContentView.swift
//  AdventOfCode2025
//
//  Created by Matt Slocum on 12/3/25.
//

import SwiftUI
import SwiftData

struct ContentView: View {
    @Environment(\.modelContext) private var modelContext
    @State private var days: [any Challenge] = []

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
//                    .navigationTitle(day.day)
                }
            }
            .navigationSplitViewColumnWidth(min: 180, ideal: 200)
        } detail: {
            Text("Select a challenge")
        }
        .onAppear {
            if days.isEmpty {
                days = ChallengeLoader.loadChallenges()
            }
        }
    }
}

#Preview {
    ContentView()
        .modelContainer(for: Item.self, inMemory: true)
}
