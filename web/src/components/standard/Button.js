import styled from 'styled-components';
import { Colors } from '../../utils/Colors';

export const InputButton = styled.input`
  padding: 5px 10px;
  cursor: pointer;
  background-color: ${Colors.GrayD};

  :hover {
    background-color: ${Colors.GrayC};
  }
`;

export const Button = styled.button`
  padding: 5px 10px;
  cursor: pointer;
  background-color: ${Colors.GrayD};

  :hover {
    background-color: ${Colors.GrayC};
  }
`;
