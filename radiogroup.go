package clui

// RadioGroup is non-interactive invisible object. It manages
// set of Radio buttons: at a time no more than one radio
// button from a group can be selected
type RadioGroup struct {
	items []*Radio
}

// CreateRadioGroup creates a new RadioGroup
func CreateRadioGroup() *RadioGroup {
	return &RadioGroup{}
}

// Selected returns the number of currently selected radio
// button inside the group or -1 if no button is selected
func (c *RadioGroup) Selected() int {
	selected := -1

	for id, item := range c.items {
		if item.Selected() {
			selected = id
			break
		}
	}

	return selected
}

// SelectItem makes the radio selected. The function returns false
// if it failed to find the radio in the radio group
func (c *RadioGroup) SelectItem(r *Radio) bool {
	found := false

	for _, item := range c.items {
		if item == r {
			found = true
			item.SetSelected(true)
		} else {
			item.SetSelected(false)
		}
	}

	return found
}

// SetSelected selects the radio by its number.
// Returns false if the number is invalid for the radio group.
func (c *RadioGroup) SetSelected(id int) bool {
	if id < 0 || id >= len(c.items) {
		return false
	}
	return c.SelectItem(c.items[id])
}

// AddItem adds a new radio button to group
func (c *RadioGroup) AddItem(r *Radio) {
	c.items = append(c.items, r)
	r.SetGroup(c)
}

// Item returns an item from the group.
// Returns nil if the number is invalid for the radio group.
func (c *RadioGroup) Item(id int) *Radio {
	if id < 0 || id >= len(c.items) {
		return nil
	}
	return c.items[id]
}

// SelectedItem returns the select item from the group.
// Returns nil if the group is empty or if none are selected.
func (c *RadioGroup) SelectedItem() *Radio {
	return c.Item(c.Selected())
}

// Active checks if one of the element is active.
func (c *RadioGroup) Active() bool {
	for _, item := range c.items {
		if item.Active() {
			return true
		}
	}

	return false
}

// SetActive sets the selected item to the given active value.
// If none are active, fallback to parent.
func (c *RadioGroup) SetActive(active bool) {
	if item := c.SelectedItem(); item != nil {
		item.SetActive(active)
		return
	}
}
