import 'dart:collection';
import 'dart:core';
import 'dart:io';

import 'package:collection/collection.dart';

enum PortalType { INSIDE, OUTSIDE }

class Point {
  int x;
  int y;

  Point(this.x, this.y);

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is Point &&
          runtimeType == other.runtimeType &&
          x == other.x &&
          y == other.y;

  @override
  int get hashCode => x.hashCode ^ y.hashCode;

  @override
  String toString() {
    return 'Point{x: $x, y: $y}';
  }

  Point operator +(Point other) {
    return Point(x + other.x, y + other.y);
  }
}

class PointOrientation {
  Point point;
  String portalCode;
  PortalType portalType;

  PointOrientation(this.point, this.portalCode, this.portalType);

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is PointOrientation &&
          runtimeType == other.runtimeType &&
          point == other.point &&
          portalType == other.portalType;

  @override
  int get hashCode => point.hashCode ^ portalType.hashCode;

  @override
  String toString() {
    return 'PointOrientation{point: $point, code: $portalCode, portalType: $portalType}';
  }
}

class Portal {
  String code;
  Point p;

  Portal(this.code, this.p);

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is Portal &&
          runtimeType == other.runtimeType &&
          code == other.code &&
          p == other.p;

  @override
  int get hashCode => code.hashCode ^ p.hashCode;

  @override
  String toString() {
    return 'Portal{code: $code, p: $p}';
  }
}

bool isUppercase(int charCode) {
  return charCode >= 65 && charCode <= 90;
}

bool isDot(int charCode) {
  return charCode == ".".codeUnitAt(0);
}

bool isOk(List<String> map, Point p, bool Function(int) okFunction) {
  try {
    if (okFunction(map[p.x].codeUnitAt(p.y))) {
      return true;
    }
  } on RangeError {
    return false;
  }
  return false;
}

Pair<Map<Point, PointOrientation>, Map<String, List<PointOrientation>>>
    findPortals(List<String> map) {
  var h = map.length;
  var w = map[0].length;
  print("${h} ${w}");

  var portalCodes = Map<String, List<PointOrientation>>();

  // find outer portals
  for (var col in [0, w - 2]) {
    for (var row = 2; row < h - 2; row++) {
      if (isUppercase(map[row].codeUnitAt(col))) {
        var code = map[row].substring(col, col + 2);
        portalCodes.putIfAbsent(code, () => []);
        if (col == 0) {
          var p = Point(row, 2);
          portalCodes[code].add(PointOrientation(p, code, PortalType.OUTSIDE));
        } else {
          var p = Point(row, w - 3);
          portalCodes[code].add(PointOrientation(p, code, PortalType.OUTSIDE));
        }
      }
    }
  }

  for (var row in [0, h - 2]) {
    for (var col = 2; col < w - 2; col++) {
      if (isUppercase(map[row].codeUnitAt(col))) {
        var code = map[row][col] + map[row + 1][col];
        portalCodes.putIfAbsent(code, () => []);
        if (row == 0) {
          var p = Point(2, col);
          portalCodes[code].add(PointOrientation(p, code, PortalType.OUTSIDE));
        } else {
          var p = Point(row - 1, col);
          portalCodes[code].add(PointOrientation(p, code, PortalType.OUTSIDE));
        }
      }
    }
  }
//  print(portalCodes);

  // find inner portals
  Point topLeft, bottomRight;

  // find topLeft
  for (var row = 2; row < h - 2 && topLeft == null; row++) {
    for (var col = 2; col < w - 2 && topLeft == null; col++) {
      if (map[row][col] == " ") {
        topLeft = Point(row, col);
      }
    }
  }
  // find bottomRight
  for (var row = h - 3; row >= 2 && bottomRight == null; row--) {
    for (var col = w - 3; col >= 2 && bottomRight == null; col--) {
      if (map[row][col] == " ") {
        bottomRight = Point(row, col);
      }
    }
  }
  print("${topLeft}   ${bottomRight}");

  // inner vertical
//  for (var col in [topLeft.y, bottomRight.y - 1]) {
//    for (var row in [topLeft.x, bottomRight.x - 1]) {
//      if (isUppercase(map[row].codeUnitAt(col))) {
//        var code = map[row][col] + map[row + 1][col];
//        portalCodes.putIfAbsent(code, () => []);
//        if (row == topLeft.x) {
//          var p = Point(row - 1, col);
//          portalCodes[code].add(p);
//        } else {
//          var p = Point(row + 2, col);
//          portalCodes[code].add(p);
//        }
//      }
//    }
//  }
//
//  // inner horizontal
//  for (var col in [topLeft.y, bottomRight.y - 1]) {
//    for (var row in [topLeft.x, bottomRight.x - 1]) {
//      if (isUppercase(map[row].codeUnitAt(col))) {
//        var code = map[row].substring(col, col + 2);
//        portalCodes.putIfAbsent(code, () => []);
//        if (row == topLeft.x) {
//          var p = Point(row - 1, col);
//          portalCodes[code].add(p);
//        } else {
//          var p = Point(row + 2, col);
//          portalCodes[code].add(p);
//        }
//      }
//    }
//  }
  for (var row = topLeft.x; row <= bottomRight.x; row++) {
    for (var col in [topLeft.y, bottomRight.y]) {
      var portal = checkPoint(map, Point(row, col));

      if (portal != null) {
        portalCodes.putIfAbsent(portal.code, () => []);

        portalCodes[portal.code]
            .add(PointOrientation(portal.p, portal.code, PortalType.INSIDE));
      }
    }
  }

  for (var row in [topLeft.x, bottomRight.x]) {
    for (var col = topLeft.y; col <= bottomRight.y; col++) {
      var portal = checkPoint(map, Point(row, col));

      if (portal != null) {
        portalCodes.putIfAbsent(portal.code, () => []);

        portalCodes[portal.code]
            .add(PointOrientation(portal.p, portal.code, PortalType.INSIDE));
      }
    }
  }

//  print(portalCodes);

  // merge the portals
  var portals = Map<Point, PointOrientation>();

  portalCodes.forEach((String code, List<PointOrientation> points) {
    if (points.length == 2) {
      var left = points[0];
      var right = points[1];
      portals[left.point] = right;
      portals[right.point] = left;
    }
  });

  return Pair(portals, portalCodes);
}

class Pair<K, V> {
  K left;
  V right;

  Pair(this.left, this.right);
}

//var directionRow = [0, 1, -1, 0];
//var directionCol = [1, 0, 0, -1];
var directionList = [Point(0, 1), Point(0, -1), Point(1, 0), Point(-1, 0)];
var directions = {
  "HORIZONTAL": [Point(0, 1), Point(0, -1)],
  "VERTICAL": [Point(1, 0), Point(-1, 0)]
};

var dotPoints = {
  "HORIZONTAL": [
    Point(0, 1),
    Point(0, -1),
  ],
  "VERTICAL": [
    Point(1, 0),
    Point(-1, 0),
  ]
};

var codeExtractors = <Point, Function>{
  Point(0, 1): (List<String> map, Point p) {
    return map[p.x].substring(p.y - 1, p.y + 1);
  },
  Point(0, -1): (List<String> map, Point p) {
    return map[p.x].substring(p.y, p.y + 2);
  },
  Point(1, 0): (List<String> map, Point p) {
    return map[p.x - 1][p.y] + map[p.x][p.y];
  },
  Point(-1, 0): (List<String> map, Point p) {
    return map[p.x][p.y] + map[p.x + 1][p.y];
  }
};

Portal checkPoint(List<String> map, Point p) {
  var x = p.x;
  var y = p.y;

  if (isOk(map, p, isUppercase)) {
//    if (isOk(map, Point(x, y + 1), isUppercase)) {
//      var code = map[x].substring(y, y + 2);
//
//      if (isOk(map, Point(x, y - 1), isDot)) {
//        return Portal(code, Point(x, y - 1));
//      }
//      if (isOk(map, Point(x, y + 2), isDot)) {
//        return Portal(code, Point(x, y + 2));
//      }
//    }

    for (var dirEntry in directions.entries) {
      var dirName = dirEntry.key;
      var dirs = dirEntry.value;

      var dotDirs = dotPoints[dirName];

      for (var dir in dirs) {
        var nextPoint = p + dir;

        if (isOk(map, nextPoint, isUppercase)) {
          for (var dotDir in dotDirs) {
            var codeExtractor = codeExtractors[dotDir];
            var code = codeExtractor(map, p);

            var dotPoint = p + dotDir;

            if (isOk(map, dotPoint, isDot)) {
              return Portal(code, dotPoint);
            }
          }
        }
      }
    }
  }
  return null;
}

int bfs(List<String> map, Map<Point, PointOrientation> portals, Point startPos,
    Point endPos) {
  var queue = Queue<Point>();
//  var visited = Set<Point>();
  var dist = Map<Point, int>();

  dist[startPos] = 0;
  queue.add(startPos);

  while (!queue.isEmpty) {
//    print(queue);

    var curPos = queue.removeFirst();
    var curDist = dist[curPos];

    for (var dir in directionList) {
      var nextPos = curPos + dir;

      var isValidPos = isOk(map, nextPos, isDot);

      if (isValidPos && !dist.containsKey(nextPos)) {
        // not visited
        dist[nextPos] = curDist + 1;
        queue.add(nextPos);
      }
    }
    // portal logic here
    var portalStep = portals[curPos];
    if (portalStep != null) {
      if (dist[portalStep.point] == null) {
        dist[portalStep.point] = curDist + 1;
        queue.add(portalStep.point);
      }
    }
  }

  return dist[endPos];
}

class State {
  Point p;
  String code;
  int level;

  State(this.p, this.code, this.level);

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is State &&
          runtimeType == other.runtimeType &&
          p == other.p &&
          level == other.level;

  @override
  int get hashCode => p.hashCode ^ level.hashCode;

  @override
  String toString() => '${code ?? ''} $p at level $level';
}

int distance(Point a, Point b) {
  return (a.x - b.x).abs() + (a.y - b.y).abs();
}

int score(State state, State finalState) {
  return (state.level + 1) * distance(state.p, finalState.p);
}

int Function(State, State) stateComparator(State finalState) {
  return (State a, State b) {
    if (a.level == b.level) {
      return distance(a.p, finalState.p).compareTo(distance(b.p, finalState.p));
    }
    return a.level.compareTo(b.level);
  };
}

int bfs2(List<String> map, Map<Point, PointOrientation> portals, Point startPos,
    Point endPos) {
  var startState = State(startPos, null, 0);
  var finalState = State(endPos, null, 0);

  var queue = HeapPriorityQueue<State>(stateComparator(finalState));
  testPriorityQeuue(queue);

  var dist = Map<State, int>();

  dist[startState] = 0;
  queue.add(startState);

  // final prev = <State, State>{};

  var steps = 1;
  while (queue.isNotEmpty && steps < 1000000 && dist[finalState] == null) {
//    print(queue);

    var curState = queue.removeFirst();
    var curDist = dist[curState];

    for (var dir in directionList) {
      var nextPos = curState.p + dir;
      var nextState = State(nextPos, null, curState.level);

      var isValidPos = isOk(map, nextPos, isDot);

      if (isValidPos && !dist.containsKey(nextState)) {
        // not visited
        dist[nextState] = curDist + 1;
        // prev[nextState] = curState;
        queue.add(nextState);
      }
    }
    // portal logic here
    var portalStep = portals[curState.p];
    if (portalStep != null) {
      State nextState;
      if (portalStep.portalType == PortalType.OUTSIDE) {
        nextState =
            State(portalStep.point, portalStep.portalCode, curState.level + 1);
      } else {
        nextState =
            State(portalStep.point, portalStep.portalCode, curState.level - 1);
      }

      if (dist[nextState] == null && nextState.level >= 0) {
        dist[nextState] = curDist + 1;
        // prev[nextState] = curState;
        queue.add(nextState);
      }
    }

    steps++;
  }

  // printPath(prev, startState, finalState);

  return dist[finalState];
}

void printPath(Map<State, State> prevMap, State startState, State endState) {
  var state = prevMap[endState];

  final path = List<State>();
  // path.add(endState);

  while (state != null && state != startState) {
    if (state == null) {
      print("state is null!");
      return;
    }
    if (state.code != null) {
      path.add(state);
    }
    state = prevMap[state];
  }

  print('Path:');
  for (final el in path.reversed) {
    print(el);
  }
}

void testPriorityQeuue(PriorityQueue<State> pq) {
  pq.add(State(Point(1, 1), null, 1));
  pq.add(State(Point(1, 1), null, 2));
  pq.add(State(Point(1, 1), null, 0));

  final bestState = pq.removeFirst();

  if (bestState.level != 0) {
    throw StateError("wrong ordering in priority queue");
  }
}

void main(List<String> args) {
  var fin = new File("input.txt");

  var fileContents = fin.readAsStringSync().split('\r\n');

  for (var line in fileContents) {
    print(line);
  }

  var tmp = findPortals(fileContents);

  var portals = tmp.left;
  var portalCodes = tmp.right;

  portals.forEach((k, v) => print("${k} => ${v}"));

  var startPoint = portalCodes["AA"][0];
  var endPoint = portalCodes["ZZ"][0];

  // print(bfs(fileContents, portals, startPoint.point, endPoint.point));
  print(bfs2(fileContents, portals, startPoint.point, endPoint.point));
}
