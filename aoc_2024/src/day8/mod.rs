use std::{
    collections::{HashMap, HashSet},
    ops::{Add, Sub},
};

struct Map {
    grid: Vec<Vec<u8>>,
    antennas: HashMap<u8, Vec<Point>>,
}

#[derive(Clone, Copy, Default, Hash, Eq)]
struct Point(i32, i32);

impl From<&str> for Map {
    fn from(input: &str) -> Self {
        let grid: Vec<Vec<u8>> = input.lines().map(|row| row.bytes().collect()).collect();
        let antennas = grid
            .iter()
            .enumerate()
            .flat_map(|(x, row)| {
                row.iter().enumerate().filter_map(move |(y, &cell)| {
                    (cell != b'.').then_some((cell, Point(x as i32, y as i32)))
                })
            })
            .fold(HashMap::new(), |mut antennas, (sym, pos)| {
                antennas.entry(sym).or_insert_with(Vec::new).push(pos);
                antennas
            });

        Self { grid, antennas }
    }
}

impl Map {
    fn signal(&self) -> HashSet<Point> {
        let mut antinodes = HashSet::new();
        let pairs = self.get_pairs();

        for (p1, p2) in pairs {
            [p1 + (p1 - p2), p2 + (p2 - p1)]
                .into_iter()
                .filter(move |&antinode| self.get(antinode).is_some())
                .for_each(|antinode| {
                    antinodes.insert(antinode);
                })
        }
        antinodes
    }

    fn harmonics(&self) -> HashSet<Point> {
        let mut antinodes = HashSet::new();
        let pairs = self.get_pairs();
        for (p1, p2) in pairs {
            [(p1, p1 - p2), (p2, p2 - p1)]
                .into_iter()
                .for_each(|(mut point, offset)| {
                    while self.get(point).is_some() {
                        antinodes.insert(point);
                        point = point + offset;
                    }
                });
        }

        antinodes
    }

    fn get_pairs(&self) -> Vec<(Point, Point)> {
        self.antennas
            .values()
            .flat_map(|antenna| {
                antenna
                    .iter()
                    .flat_map(|&p1| antenna.iter().map(move |&p2| (p1, p2)))
                    .filter(|(p1, p2)| p1 != p2)
            })
            .collect()
    }

    fn get(&self, Point(x, y): Point) -> Option<u8> {
        self.grid.get(x as usize)?.get(y as usize).copied()
    }
}

impl PartialEq for Point {
    fn eq(&self, other: &Self) -> bool {
        self.0 == other.0 && self.1 == other.1
    }
}

impl Add for Point {
    type Output = Self;

    fn add(self, other: Self) -> Self {
        Self(self.0 + other.0, self.1 + other.1)
    }
}

impl Sub for Point {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        Self(self.0 - other.0, self.1 - other.1)
    }
}

pub fn resonant_collinearity(input: &str) -> u32 {
    let map = Map::from(input);
    let count = map.signal().len() as u32;

    count
}

pub fn resonant_collinearity_2(input: &str) -> u32 {
    let map = Map::from(input);
    let count = map.harmonics().len() as u32;

    count
}
