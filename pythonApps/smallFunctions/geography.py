"""
>>> from math import isclose

Comiskey Park
>>> comiskey = Coordinate(41.831141, -87.6361152, input_units="degrees")
>>> rad = comiskey.astuple()
>>> isclose(rad[0], 0.7300911, abs_tol=1e-4)
True
>>> isclose(rad[1], -1.529539, abs_tol=1e-4)
True

The Astrodome
>>> astrodome = Coordinate(29.6850162,-95.4099542, input_units="degrees")
>>> isclose(comiskey.distance(astrodome), 945.174209, abs_tol=1e-4)
True

Wrigley Field
>>> wrigley = Coordinate(0.7321384359631922, -1.5299124550698397)
>>> deg = wrigley.astuple("degrees")
>>> isclose(deg[0], 41.948442, abs_tol=1e-4)
True
>>> isclose(deg[1], -87.657527, abs_tol=1e-4)
True
>>> isclose(comiskey.distance(wrigley), 8.18386, abs_tol=1e-4)
True

"""
from math import asin, sqrt, sin, cos, radians, isclose

class Coordinate:
    def __init__(self,lattitude,longitude, input_units="radians"):
        if lattitude == None:
            lattitude = input("please input a lattitude in either radians or degrees: ")
        if longitude == None:
            longitude = input("please input a longitude in either radians or degrees: ")
        #if the user degrees for the input_unit then they do not need to be prompted to input a unit
        if input_units == "radians":
            input_units = input("your current input is unit is radians, pleasae type to y to keep or n for degrees: ")
        valid_inputs = ["radians","degrees","y","n"]
        while input_units not in valid_inputs:
            input_units = input("please input a valid input of radians or degrees")
        if input_units.lower() == 'y':
            input_units = "radians"
        elif input_units.lower() == 'n':
            input_units = "degrees"
        if input_units == "degrees":
            lattitude = radians(lattitude)
            longitude = radians(longitude)
            input_units = "radians"
        self.lat = lattitude
        self.long = longitude
        self.input_units = input_units


    def __repr__(self):
        return f"{self.__class__.__name__}({self.lat}, {self.long})"
        

    
    """ Represents a (latitude, longitude) coordinate on Earth

    The user can pass the arguments :data:`latitude` and :data:`longitude` as either radians or degrees, as specified by the argument :data:`input_units`.
    However, the attributes :data:`at` and :data:`long` must be represented in radians.

    Parameters
    ----------
    latitude : float
        Latitude in units of :data:`input_units`
    longitude : float
        Longitude in units of :data:`input_units`
    input_units : str, optional
        The units of the parameters :data:`latitude` and :data:`longitude`
        Valid values are ``"radians"`` or ``"degrees"``. Defaults to ``"radians"``

    Attributes
    ----------
    lat : float
        Latitude in radians
    long : float
        Longitude in radians
    """


    def astuple(self, units="radians"):
        """ Returns a tuple of coordinates ``(lat, long)`` with specified units

        Parameters
        ----------
        units : str, optional
            Valid values are ``"radians"`` or ``"degrees"``. Defaults to ``"radians"``


        Returns
        -------
        tuple[float]:
            A tuple of ``(lat, long)`` with the specified units
        """
        if units == "radians":
            return((self.lat, self.long))
        elif units == "degrees":
            return((radians(self.lat), radians(self.long)))
        else:
            print("you have inputted an invalid unit")


    def distance(self, coord):
        """ Returns the distance between this and another coordinate.

        The distance between the coordinates should be calculated from the Haversine formula:

        .. math::

            d = 2r \\arcsin{\\left(\\sqrt{\\sin^2{\\left(\\frac{\\varphi_2 - \\varphi_1}{2}\\right) + \\cos{\\varphi_1}\\cos{\\varphi_2}\\sin^2{\\left(\\frac{\\lambda_2 - \\lambda_1}{2}\\right)}}}\\right)}

        where :math:`\\left(\\varphi_1, \\lambda_1\\right)`  and :math:`\\left(\\varphi_2, \\lambda_2\\right)`
        are the coordinates in radians; and  :math:`r = 3961` is the radius of the Earth in miles.

        Parameters
        ----------
        coord : Coordinate
            A second coordinate

        Returns
        -------
        float:
            The distance between this coordinate instance and the other :data:`coord`


        """
        r = 3961   
        coord = str(coord)
        coord = coord.replace('(','')
        coord = coord.replace(')','')
        coord = coord.split(',')
        d = 2*r*asin(sqrt(sin((float(coord[0]) - self.lat)/2)**2 + cos(self.lat)*cos(float(coord[0]))*sin((float(coord[1]) - self.long)/2)**2))
        return(d)

    def show_map(self):
        """ Shows this coordinate on Google Maps in a browser window.

        This can be accomplished with the Python standard library function :py:func:`webbrowser.open_new_tab`.
        The URL you need is ``http://maps.google.com/maps?q=<latitude>,<latitude>`` where ``<latitude>`` and
        ``<longitude>`` are the coordinates in degrees.

        For example, given that the Chicago Bean's street address is at located at the coordinates
        (41.8821512, -87.6246838), then we can show it on Google Maps at ::

            import webbrowser
            webbrowser.open_new_tab(r'http://maps.google.com/maps?q=41.8821512,-87.6246838')

        """
        import webbrowser
        url = "http://maps.google.com/maps?q={},{}".format(self.lat*180/math.pi, self.long*180/math.pi)
        webbrowser.open_new_tab(url)




    def __str__(self):
        return f"({self.lat:0.7}, {self.long:0.7})"


if __name__ == "__main__":
    # The locations used in doctests.  Here for your convenience
    comiskey = Coordinate(41.831141, -87.6361152, input_units="degrees")
    wrigley = Coordinate(0.7321384359631922, -1.5299124550698397)
    astrodome = Coordinate(29.6850162, -95.4099542, input_units="degrees")
    print(astrodome, 'astrodome')
    comtup = comiskey.astuple("radians")
    print(comtup,'comiskey') 
    print(comiskey)
    print(isclose(comtup[0], 0.7300911, abs_tol=1e-4),'1')
    print(isclose(comtup[1], -1.5299124550698397, abs_tol=1e-3),'2')
    print(isclose(comiskey.distance(wrigley), 8.18386, abs_tol=1e-3),'3')
    

    #locdist = location.distance()
    #locmap = location.show_map()