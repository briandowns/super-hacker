package ruby

// http://sandbox.mc.edu/~bennet/ruby/code/blink_rb.html

const ReflexTest = `
#!/usr/bin/ruby

# Import the library.
require 'tk'

# Parameters.
Width = 5               # Width of button grid.
Height = 5              # Height of button grid.
MinWait = 200           # Smallest button change wait (ms)
MaxWait = 1400          # Largest button change wait (ms)
InitWait = 800          # Initial button change wait (ms)
LossRate = 2000         # Frequency to take away points.

# Set defaults.  Some we keep in constants to use later.
BG = '#ccffcc'
TkOption.add('*background', BG)
TkOption.add('*activeBackground', '#ddffdd')
FG = '#006600'
TkOption.add('*foreground', FG)
TkOption.add('*activeForeground', FG)
TkOption.add('*troughColor', '#99dd99')

# Root window.
root = TkRoot.new('background' => BG) { title 'Click Fast' }

# Button from the panel
class PanelButton < TkButton
private
  # Exchange colors on the button.
  def cswap
    for p in [['background', 'foreground'], 
        ['activebackground', 'activeforeground']]
      c = cget(p[0])
      configure(p[0] => cget(p[1]))
      configure(p[1] => c)
    end
  end
public
  # Initialize the button within the widget sup, at position pos (zero-based)
  # with the number num.  When pressed, send the score (+ or -) to cmd.
  # Scorekeeper is an object which implements an up and down methods to
  # receive score changes.
  def initialize(sup, pos, num, scorekeeper)
    super(sup, 'text' => num.to_s, 'command' => proc { self.pushed },
          'activeforeground' => '#990000', 'activebackground' => '#ffdddd')
    grid('row' => pos / Width + 1, 'column' => pos % Width, 'sticky' => 'news')
    @active = false
    @scorekeeper = scorekeeper
  end
  attr_reader :active

  # Activate or deactivate the button.
  def activate
    if not @active
      cswap
      @active = true
    end
  end
  def deactivate
    if @active
      cswap
      @active = false
    end
  end

  # When pushed, send our number, or negative our number, to the scorekeeping
  # command.
  def pushed
    n = self.cget('text').to_i
    if @active
      @scorekeeper.up(n)
    else
      @scorekeeper.down(n)
    end
  end
end

# This class calls reduces the score at the indicated time rate.  
class ScoreTimer
  # This object will call scorekeeper.down(step) each rate ms.
  def initialize(scorekeeper, rate = 500, step = 1)
    @scorekeeper = scorekeeper
    @rate = rate
    @step = step

    Tk.after(rate, proc { self.change })    
  end

  # Reduce the score periodically
  def change
    @scorekeeper.down(@step)
    Tk.after(@rate, proc { self.change })
  end
end

# This is a box displaying a count-up timer in minutes and seconds to tenths
# m:ss.d
class TimeCounter < TkLabel
  # Initialize.  Displays zero and starts the ticking event.
  def initialize(root)
    super(root, "text" => '0:00.0', 'anchor' => 'e')
    @count = 0
    Tk.after(100, proc { self.change })
  end

  # One clock tick (tenths of a second).  Increment the counter, then build
  # the new display value.
  def change
    @count += 1
    self.configure('text' => 
                     sprintf("%d:%02d.%d", 
                             @count / 600, (@count / 10) % 60, @count % 10))
    Tk.after(100, proc { self.change })
  end
end

# This is the main application GUI.
class App
private
  # Set the score value.
  def setscore(val)
    color = if val < 0 then 'red' else FG end
    @slab.configure('text' => val.to_s, 'foreground' => color)
  end
public
  # The wait attribute is the amount of time (ms) between button changes.
  attr_writer :wait

  # Initialize it and have the applicate drawn in the root window.
  def initialize(root)
    # This is the label containing the score.  Initially zero.
    @slab = TkLabel.new(root) {
      text "0"
      anchor 'e'
      grid('row' => 0, 'column' => 0, 'columnspan' => Width / 2, 
           'sticky' => 'w')
    }

    # This is the timer window at upper right.
    TimeCounter.new(root).
      grid('row' => 0, 'column' => Width/2, 'columnspan' => (Width+1)/2, 
           'sticky' => 'e')

    # Create the buttons.  First, make an array of numbers from 1 to the
    # number of buttons, then create the buttons, each labelled with a 
    # number chosen at random from the list, so thare are no repeats.
    nums = (1..Height*Width).to_a;
    @buts= [ ]
    for n in (0...Height*Width)
      pos = rand(nums.length)
      @buts.push(PanelButton.new(root, n, nums[pos], self))
      nums.delete_at(pos)
    end

    # This creates the slider to adjust the speed of the game.  The proc is
    # called whenever the slider changes, and is sent the new setting.
    scale = TkScale.new('command' => proc { |v| self.wait = v.to_i } ) {
      orient "horizontal"       # Which way the slider goes.
      from MinWait              # Value of smallest setting
      to MaxWait                # Value of largest setting
      showvalue false           # Don't show the numeric value of the setting.
      grid('row' => Height + 1, 'column' => 1, 'columnspan' => Width - 2,
           'sticky' => 'news')
    }
    scale.set(InitWait)

    # Labels by the slider.
    TkLabel.new {
      text "Fast"
      anchor "w"
      grid("row" => Height + 1, 'column' => 0, 'sticky' => 'w')
    }
    TkLabel.new {
      text "Slow"
      anchor "e"
      grid("row" => Height + 1, 'column' => Width-1, 'sticky' => 'e')
    }

    @wait = InitWait

    # Decrement the score every LossRate period.
    @timer = ScoreTimer.new(self, LossRate)
    self.change
  end

  # Actions to increase or decrease the score.
  def up(delta)
    setscore(@slab.cget('text').to_i + delta)
  end
  def down(delta)
    setscore(@slab.cget('text').to_i - delta)
  end

  # Change (or set, if none is yet set) the active button.  It deactivates
  # the button in @buts[0], It then chooses some other button at random,
  # activates that and swaps it into position 0.
  def change
    @buts[0].deactivate
    pos = rand(@buts.length - 1) + 1
    @buts[0], @buts[pos] = @buts[pos], @buts[0]
    @buts[0].activate

    Tk.after(@wait, proc { self.change })
  end
end

a = App.new(root)

Tk.mainloop
`
