#!/usr/bin/env ruby
# encoding: UTF-8

require 'forwardable'
require 'securerandom'

def rand_num(n)
  SecureRandom.random_number(n)
end

STR_LEN = 100.freeze

class Pool
  CAPACITY = 100.freeze
  attr_accessor :pool

  def initialize
    @pool = []
  end

  def eliminate(&block)
    count = -1
    @pool.reject! do |x|
      res = block.call(x)
      count += 1 if res
      res && count < 100
    end
  end

  def inspect
    "Pool: (size: #{size})\n" +
      @pool.map do |x|
        "#{x.inspect}: #{x.count('1')}"
      end.join("\n")
  end

  extend Forwardable

  def_delegators :@pool, :push, :size, :at, :sort_by!
end

class Evolutation
  class << self
    TIME = 300

    def evolutation(pool)
      seed pool
      eli_fiber = Fiber.new do
        loop do
          eliminate pool
          Fiber.yield
        end
      end

      mul_fiber = Fiber.new do
        loop do
          multiplication pool
          eli_fiber.resume
          Fiber.yield
        end
      end

      TIME.times do
        mul_fiber.resume
      end

      p pool
    end

    def seed(pool)
      pool.push random_dna
      pool.push random_dna
    end

    def multiplication(pool)
      target = (Pool::CAPACITY * 2).freeze
      while pool.size < target
        parent1, parent2 = pick_parent pool
        child1, child2, = cross parent1, parent2
        child1 = mutation child1
        child2 = mutation child2
        pool.push child1
        pool.push child2
      end
    end

    def eliminate(pool)
      pool.pool = pool.sort_by! { |d| d.count('1') }.reverse.take Pool::CAPACITY
    end

    private

    def random_dna
      STR_LEN.times.map { [0, 1].sample }.join
    end

    def reverse(ch)
      ch == '1' ? '0' : '1'
    end

    def pick_parent(pool)
      size = pool.size
      p1 = rand_num size
      p2 = rand_num size
      [pool.at(p1), pool.at(p2)]
    end

    def cross(parent_a, parent_b)
      n = rand_num STR_LEN
      a_chop = [parent_a.slice(0...n), parent_a.slice(n..-1)]
      b_chop = [parent_b.slice(0...n), parent_b.slice(n..-1)]
      [a_chop[0] + b_chop[1], b_chop[0] + a_chop[1]]
    end

    def mutation(dna)
      dna.dup.tap do |d|
        n = rand_num STR_LEN
        n.times do
          idx = rand_num STR_LEN
          d[idx] = reverse d[idx]
        end
      end
    end
  end
end

def main
  pool = Pool.new
  Evolutation.evolutation pool
end

main if __FILE__ == $PROGRAM_NAME
