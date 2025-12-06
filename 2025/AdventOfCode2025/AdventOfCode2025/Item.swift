//
//  Item.swift
//  AdventOfCode2025
//
//  Created by Matt Slocum on 12/3/25.
//

import Foundation
import SwiftData

@Model
final class Item {
    var timestamp: Date
    
    init(timestamp: Date) {
        self.timestamp = timestamp
    }
}
