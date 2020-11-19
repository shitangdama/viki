require("@rails/ujs").start()
require("turbolinks").start()
require("@rails/activestorage").start()
require("channels")

import 'bootstrap/dist/js/bootstrap'
import "@fortawesome/fontawesome-free/js/all";

import './volt'

import 'styles/admin'

// Uncomment to copy all static images under ../images to the output folder and reference
// them with the image_pack_tag helper in views (e.g <%= image_pack_tag 'rails.png' %>)
// or the `imagePath` JavaScript helper below.
//

const images = require.context('../images', true)
const imagePath = (name) => images(name, true)