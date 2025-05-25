function Prompt() {
    let toast = function (c) {
        const {
            msg = '',
            icon = 'success',
            position = 'top-end',

        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })

    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })

    }

    async function custom(c) {
        const {
            icon = "",
            msg = "",
            title = "",
            showConfirmButton = true,
        } = c;

        const {value: result} = await Swal.fire({
            icon: icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton: showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result);
                    }
                } else {
                    c.callback(false);
                }
            } else {
                c.callback(false);
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

function sayHi() {
    alert("Hi!");
}


function chooseYourDates(roomNum, csrfToken) {
    const checkButton = document.getElementById("check-availability-button");

    if (!checkButton) {
        console.error("Button with ID 'check-availability-button' not found.");
        return;
    }

    checkButton.addEventListener("click", function () {
        let html = `
        <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
            <div class="form-row">
                <div class="col">
                    <div class="form-row" id="reservation-dates-modal">
                        <div class="col">
                            <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
                        </div>
                        <div class="col">
                            <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
                        </div>
                    </div>
                </div>
            </div>
        </form>
        `;

        // Ensure 'attention' is available. It might be initialized globally in app.js
        // or passed around, or initialized in the template before this function is called.
        // For this example, we assume 'attention' is globally accessible from app.js.
        if (typeof attention === 'undefined') {
            console.error("The 'attention' object (from Prompt) is not defined. Make sure it's initialized before calling chooseYourDates.");
            return;
        }

        attention.custom({
            title: 'Choose your dates',
            msg: html,
            willOpen: () => {
                const elem = document.getElementById("reservation-dates-modal");
                if (elem) {
                    const rp = new DateRangePicker(elem, {
                        format: 'yyyy-mm-dd',
                        showOnFocus: true,
                        minDate: new Date(),
                    });
                } else {
                    console.error("Element 'reservation-dates-modal' not found for DateRangePicker.");
                }
            },
            didOpen: () => {
                const startInput = document.getElementById("start");
                const endInput = document.getElementById("end");
                if (startInput) startInput.removeAttribute("disabled");
                if (endInput) endInput.removeAttribute("disabled");
            },
            callback: function(result) {
                console.log("Modal callback called with result:", result);

                // The 'result' from Swal.fire in your 'custom' function is an array [startDate, endDate]
                // or an object with 'dismiss' if cancelled.
                // You might want to check if the modal was confirmed before proceeding.
                if (!result || result.dismiss) {
                    console.log("Modal was dismissed or cancelled.");
                    return;
                }

                let form = document.getElementById("check-availability-form");
                if (!form) {
                    console.error("Form 'check-availability-form' not found in modal.");
                    return;
                }

                let formData = new FormData(form);
                formData.append("csrf_token", csrfToken); // Use the csrfToken parameter
                formData.append("room_id", roomNum);      // Use the roomNum parameter

                fetch('/search-availability-json', {
                    method: "post",
                    body: formData,
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.ok) {
                            attention.custom({
                                icon: 'success',
                                showConfirmButton: false,
                                msg: '<p>Room is available!</p>'
                                    + '<p><a href="/book-room?id='
                                    + data.room_id
                                    + '&s='
                                    + data.start_date
                                    + '&e='
                                    + data.end_date
                                    + '" class="btn btn-primary">' // Added space before class
                                    + 'Book now!</a></p>',       // Corrected closing </a>
                            });
                        } else {
                            attention.error({
                                msg: data.message || "No availability", // Use server's message if available
                            });
                        }
                    })
                    .catch(error => {
                        console.error("Error fetching availability:", error);
                        attention.error({
                            msg: "An error occurred while checking availability.",
                        });
                    });
            }
        });
    });
}