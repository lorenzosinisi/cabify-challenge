require 'minitest/unit'
require 'minitest/autorun'

require 'cabify_challenge'

module CabifyChallenge
  class PricingTest < Minitest::Unit::TestCase

    def test_default_price
      price = Pricing::DefaultPrice.new(1.95)

      assert_equal 0.0, price.calculate(0)
      assert_equal 1.95, price.calculate(1)
      assert_equal 3.90, price.calculate(2)
    end
  end
end
